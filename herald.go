package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/c0d3x-io/herald/internal/config"
	"github.com/c0d3x-io/herald/internal/proxy"
)

func main() {
	// Load log ---> YAGNI — You Aren't Gonna Need It.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("config error", "error", err)
		os.Exit(1)
	}
	logger.Info("---Herald Started--",
		"ListenerAddr", cfg.ListenAddr,
		"UpstreamURL", cfg.UpstreamURL,
	)

	// Proxy sever to ensure it's working and online
	server := http.NewServeMux()
	server.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Herald:ok")) // ignore error
	})

	// Proxy
	herald, err := proxy.New(cfg.UpstreamURL, cfg.CABundlePath, logger)
	if err != nil {
		logger.Error("failed to create proxy", "error", err)
		os.Exit(1)
	}

	server.Handle("/", herald)

	// Timeouts - Security Practices
	srv := &http.Server{
		Addr:         cfg.ListenAddr,
		Handler:      server,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  90 * time.Second,
	}

	logger.Info("listening", "addr", cfg.ListenAddr)
	if err := srv.ListenAndServeTLS(cfg.TLSCert, cfg.TLSKey); err != nil {
		logger.Error("server failed", "error", err)
		os.Exit(1)
	}
}
