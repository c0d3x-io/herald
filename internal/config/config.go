// Config holds all Herald reuntime configuration.
// All values are drive by enviormental variable - no config files
package config

import (
	"fmt"
	"os"
)

type Config struct {
	ListenAddr  string // Herald binds to inside the mesh network
	UpstreamURL string // Required. UpstreamURL is the enterprise aress gateway Herald forwards to.
	LogLevel    string // Values : debug | info | warn | error
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func LoadConfig() (*Config, error) {

	upstream := os.Getenv("HERALD_UPSTREAM_URL")
	if upstream == "" {
		return nil, fmt.Errorf("HERALD_UPSTREAM_URL is required")
	}

	return &Config{
		ListenAddr:  getEnv("HERALD_LISTEN_ADDR", ":8080"),
		UpstreamURL: upstream,
		LogLevel:    getEnv("HERALD_LOG_LEVEL", "info"),
	}, nil
}
