// Package lib - this parts contains the *http.Client used to make requests.
//
// Could possibly just use another HTTP library that has some of these constants defined ...
package lib

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// From https://golang.org/pkg/net/http/#pkg-overview:
//
// For control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:
//
// tr := &http.Transport{
// 	MaxIdleConns:       10,
// 	IdleConnTimeout:    30 * time.Second,
// 	DisableCompression: true,
// }
// client := &http.Client{Transport: tr}
// resp, err := client.Get("https://example.com")
const (
	HTTPClientTimeoutDefault = 15 * time.Second
	MaxIdleConns             = 10
	IdleConnTimeout          = 30 * time.Second
	DisableCompression       = true
)

// HTTPClient - Just wraps *http.Client.
type HTTPClient struct {
	*http.Client `json:"net_http_client,omitempty"`
	LoggerParent *zerolog.Logger `json:"logger_parent,omitempty"` // Logger: parent logger. Use `Crawler.Logger()` to get contextual logger.
}

// NewHTTPClient - Make new client for issuing requests.
func NewHTTPClient(insecureSkipVerify bool, timeout time.Duration, loggerParent *zerolog.Logger) (*HTTPClient, error) {
	logger := loggerParent.
		With().Str("func", "NewHTTPClient").
		Interface("insecureSkipVerify", insecureSkipVerify).
		Interface("timeout", timeout).
		Logger()

	pool, err := x509.SystemCertPool()
	if err != nil {
		errWrapped := errors.Wrapf(
			err,
			"could not create HTTPClient due to failure in x509.SystemCertPool() - pool=%v",
			pool,
		)
		logger.Error().Err(errWrapped).Interface("pool", pool).Msg("could not create HTTPClient")

		return nil, errWrapped
	}

	// G402: TLS InsecureSkipVerify set true. (gosec)go-lint
	// nolint:gosec
	tlsConfig := tls.Config{
		RootCAs:            pool,
		InsecureSkipVerify: insecureSkipVerify,
		Renegotiation:      tls.RenegotiateFreelyAsClient, // Turn off is this is an issue.
	}
	tlsConfig.BuildNameToCertificate()

	// Does _not_ print outs the requests and responses
	transportDefault := &http.Transport{
		TLSClientConfig:    &tlsConfig,
		MaxIdleConns:       MaxIdleConns,
		IdleConnTimeout:    IdleConnTimeout,
		DisableCompression: DisableCompression,
	}
	// Prints outs the requests and responses
	transportDebug := &DebugTransport{
		Transport:    transportDefault,
		LoggerParent: loggerParent,
	}
	_, _ = transportDebug, transportDefault

	transport := transportDefault

	httpClient := &HTTPClient{
		Client: &http.Client{
			Transport: transport,
			Timeout:   timeout,
		},
		LoggerParent: loggerParent,
	}

	return httpClient, nil
}

// NewRequest -  request.
func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return SetDefaultHeaders(req), nil
}
