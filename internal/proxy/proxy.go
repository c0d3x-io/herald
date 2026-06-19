package proxy

import (
	"crypto/tls"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Herald struct {
	rp       *httputil.ReverseProxy
	logger   *slog.Logger
	upstream *url.URL
}

// New creates a Herald proxy that forwards all traffic to upstreamURL.
func New(upstreamURL string, logger *slog.Logger) (*Herald, error) {
	targetURL, err := url.Parse(upstreamURL)
	if err != nil {
		return nil, err
	}

	// Core fix: Initialize the ReverseProxy directly.
	// This avoids setting the implicit legacy Director field, preventing the runtime panic.
	rp := &httputil.ReverseProxy{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Rewrite: func(pr *httputil.ProxyRequest) {
			// SetURL safely routes the request to your target upstream.
			// It updates the Scheme, Host, Path, and the outgoing Host header automatically.
			pr.SetURL(targetURL)

			// Set standard tracking headers (X-Forwarded-For, X-Forwarded-Host, X-Forwarded-Proto)
			pr.SetXForwarded()
		},
	}

	return &Herald{
		upstream: targetURL,
		rp:       rp,
		logger:   logger,
	}, nil
}

// ServeHTTP handles incoming requests, logs them, and forwards them upstream.
func (h *Herald) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Proxying request",
		"method", r.Method,
		"path", r.URL.Path,
		"remote", r.RemoteAddr,
		"upstream", h.upstream.Host,
	)

	// Passes the client connection over to the reverse proxy engine for execution.
	h.rp.ServeHTTP(w, r)
}
