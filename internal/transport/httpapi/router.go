package httpapi

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"example.com/office-template/internal/modules/tasks"
	"example.com/office-template/internal/platform/httpx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependencies struct {
	DB       *pgxpool.Pool
	Tasks    *tasks.Service
	DemoMode bool
	Logger   *slog.Logger
}

func NewRouter(d Dependencies) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := d.DB.Ping(ctx); err != nil {
			httpx.WriteError(w, http.StatusServiceUnavailable, "NOT_READY", "database unavailable")
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "ready"})
	})
	if d.DemoMode {
		h := taskHandler{service: d.Tasks}
		mux.HandleFunc("GET /api/v1/tasks", h.list)
		mux.HandleFunc("POST /api/v1/tasks", h.create)
		mux.HandleFunc("PATCH /api/v1/tasks/{id}", h.update)
		mux.HandleFunc("DELETE /api/v1/tasks/{id}", h.delete)
	}
	return withSecurityHeaders(withAccessLog(mux, d.Logger))
}
func withSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("X-Frame-Options", "DENY")
		next.ServeHTTP(w, r)
	})
}
func withAccessLog(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		next.ServeHTTP(w, r)
		// Never log request bodies, cookies, authorization headers or query strings.
		logger.Info("http request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(started).String())
	})
}
