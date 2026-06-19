package main

import (
	"log/slog"
	"net/http"
	"os"

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
		w.Write([]byte("Herald:ok"))
	})

	// Proxy
	herald, err := proxy.New(cfg.UpstreamURL, logger)
	if err != nil {
		logger.Error("failed to create proxy", "error", err)
		os.Exit(1)
	}

	server.Handle("/", herald)
	http.ListenAndServeTLS(":8080", "localhost+2.pem", "localhost+2-key.pem", server)
}
