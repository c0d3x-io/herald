// Config holds all Herald reuntime configuration.
// All values are drive by enviormental variable - no config files
package config

import (
	"fmt"
	"os"
)

type Config struct {
	ListenAddr   string // Herald binds to inside the mesh network
	UpstreamURL  string // Required. UpstreamURL is the enterprise aress gateway Herald forwards to.
	LogLevel     string // Values : debug | info | warn | error
	CABundlePath string // Where the self sign cert are or org cert will be.
	TLSCert      string // tls1.2 certificates
	TLSKey       string // tls1.2 key
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
		ListenAddr:   getEnv("HERALD_LISTEN_ADDR", ":8080"),
		UpstreamURL:  upstream,
		LogLevel:     getEnv("HERALD_LOG_LEVEL", "info"),
		CABundlePath: getEnv("HERALD_CA_BUNBLE", ""),
		TLSCert:      getEnv("HERALD_TLS_Cert", "CaBundle/localhost+2.pem"),
		TLSKey:       getEnv("HERALD_TLS_KEY", "CaBundle/localhost+2-key.pem"),
	}, nil
}
