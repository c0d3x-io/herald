package cert

import (
	"crypto/x509"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// loadCertPoolFromDir reads every .pem/ .crt file in dir into  a cert pool
func loadCertPoolFromDir(dir string) (*x509.CertPool, error) {
	pool := x509.NewCertPool()
	count := 0

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read cert dir %q: %w", dir, err)
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(e.Name()))
		if ext != ".pem" && ext != ".crt" {
			continue
		}

		path := filepath.Join(dir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read cert %q: %w", path, err)
		}
		if pool.AppendCertsFromPEM(data) {
			count++
		}
	}
	if count == 0 {
		return nil, fmt.Errorf("no valid certificates found in %q", dir)
	}
	return pool, nil
}

// loadCertPoolFromFile
func loadCertPoolFromFile(path string) (*x509.CertPool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read CA bundle %q: %w", path, err)
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(data) {
		return nil, fmt.Errorf("no vaild certificates in %q", path)
	}
	return pool, nil
}

// Load cert for self signed or enterprise Certification
func LoadCertPool(path string) (*x509.CertPool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("Stat CA path %q: %w", path, err)
	}

	if info.IsDir() {
		return loadCertPoolFromDir(path)
	}
	return loadCertPoolFromFile(path)
}
