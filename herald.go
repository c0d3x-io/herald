package main

import (
	"github.com/c0d3x-io/herald/internal/config"
	"log/slog"
	"net/http"
	"os"
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

	http.ListenAndServe(":8080", server)

}
