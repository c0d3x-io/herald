package main

import (
	"fmt"
	"github.com/c0d3x-io/herald/internal/config"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("config error", "error", err)
		os.Exit(1)
	}
	fmt.Println("---Herald Started--")
	fmt.Printf("ListerAddr: %s", cfg.ListenAddr)
	fmt.Printf("Upstream: %s", cfg.UpstreamURL)
}
