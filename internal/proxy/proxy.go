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

// creates a Herald proxy that forwards all traffic to upstreamURL.
func New(upstreamURL string, logger *slog.Logger) (*Herald, error) {
	targetURL, err := url.Parse(upstreamURL)
	if err != nil {
		return nil, err
	}

	rp := httputil.NewSingleHostReverseProxy(targetURL)

	rp.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	rp.Rewrite = func(pr *httputil.ProxyRequest) {
		// Capture the exact domain Python is trying to reach (e.g., oauth2.googleapis.com)
		targetHost := pr.In.URL.Host
		targetScheme := pr.In.URL.Scheme

		if targetScheme == "" {
			targetScheme = "https" // Default fallback for secure traffic
		}

		// Dynamically rewrite outbound target parameters on the fly
		pr.Out.URL.Scheme = targetScheme
		pr.Out.URL.Host = targetHost
		pr.Out.Host = targetHost

		// Set standard upstream tracking headers
		pr.SetXForwarded()
	}

	return &Herald{
		upstream: targetURL,
		rp:       rp,
		logger:   logger,
	}, nil

}

// ServeHTTP handles incoming requests.
// V1: logs the request and proxies it upstream.
func (h *Herald) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Proxying request",
		"method", r.Method,
		"path", r.URL.Path,
		"remote", r.RemoteAddr,
		"upstream", h.upstream.Host,
	)
}
