package main

import (
	"github.com/c0d3x-io/herald/internal/config"
	"log/slog"
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
}
