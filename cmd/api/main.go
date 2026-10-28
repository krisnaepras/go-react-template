package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/office-template/internal/modules/tasks"
	"example.com/office-template/internal/platform/config"
	"example.com/office-template/internal/platform/database"
	"example.com/office-template/internal/transport/httpapi"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg, err := config.Load()
	if err != nil {
		logger.Error("invalid configuration", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	pool, err := database.Open(ctx, cfg.DatabaseURL)
	cancel()
	if err != nil {
		logger.Error("database not available", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	svc := tasks.NewService(tasks.NewPostgresRepository(pool))
	handler := httpapi.NewRouter(httpapi.Dependencies{
		DB: pool, Tasks: svc, DemoMode: cfg.DemoMode, Logger: logger,
	})
	srv := &http.Server{
		Addr: cfg.HTTPAddr, Handler: handler,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
		// WriteTimeout is intentionally omitted: add a per-route policy before SSE.
	}

	stopCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-stopCtx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			logger.Error("shutdown failed", "error", err)
		}
	}()

	logger.Info("API starting", "addr", cfg.HTTPAddr, "demo_mode", cfg.DemoMode)
	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}
