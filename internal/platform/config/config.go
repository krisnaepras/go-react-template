package config

import (
	"errors"
	"net"
	"os"
	"strings"
)

type Config struct {
	HTTPAddr    string
	DatabaseURL string
	DemoMode    bool
}

func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:    getDefault("HTTP_ADDR", "127.0.0.1:8080"),
		DatabaseURL: strings.TrimSpace(os.Getenv("DATABASE_URL")),
		DemoMode:    strings.EqualFold(os.Getenv("DEMO_MODE"), "true"),
	}
	if cfg.DatabaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}
	if cfg.DemoMode {
		if os.Getenv("APP_ENV") != "development" {
			return Config{}, errors.New("DEMO_MODE requires APP_ENV=development")
		}
		host, _, err := net.SplitHostPort(cfg.HTTPAddr)
		if err != nil || (host != "localhost" && host != "127.0.0.1" && host != "::1") {
			return Config{}, errors.New("DEMO_MODE requires HTTP_ADDR bound to loopback only")
		}
	}
	return cfg, nil
}
func getDefault(k, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(k)); v != "" {
		return v
	}
	return fallback
}
