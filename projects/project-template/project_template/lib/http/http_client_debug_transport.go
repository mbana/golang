// Package lib - This is a debug.
package lib

import (
	"net/http"
	"net/http/httputil"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// DebugTransport - logs the requests and responses.
type DebugTransport struct {
	*http.Transport
	LoggerParent *zerolog.Logger
}

func (d *DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	logger := d.Logger().With().Str("func", "RoundTrip").Interface("req", req).Logger()

	SetDefaultHeaders(req)

	requestBytes, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		errWrapped := errors.Wrapf(
			err,
			"failed in DebugTransport.RoundTrip due to httputil.DumpRequestOut - requestBytes=%v",
			requestBytes,
		)
		logger.Error().Err(errWrapped).Interface("requestBytes", requestBytes).Msg("could not do httputil.DumpRequestOut(req, true)")

		return nil, errWrapped
	}

	// safe to reassign a new logger with the requestBytes
	logger = logger.With().Interface("requestBytes", requestBytes).Logger()

	res, err := d.Transport.RoundTrip(req)
	if err != nil {
		errWrapped := errors.Wrapf(
			err,
			"failed in DebugTransport.RoundTrip due to d.Transport.RoundTrip(req) - res=%v",
			res,
		)
		logger.Error().Err(errWrapped).Interface("res", res).Msg("could not do d.Transport.RoundTrip(req)")

		return nil, errWrapped
	}

	// safe to reassign a new logger with the res
	logger = logger.With().Interface("res", res).Logger()

	responseBytes, err := httputil.DumpResponse(res, true)
	if err != nil {
		errWrapped := errors.Wrapf(
			err,
			"failed in DebugTransport.RoundTrip due to httputil.DumpResponse(res, true) - responseBytes=%v",
			responseBytes,
		)
		logger.Error().Err(errWrapped).Interface("responseBytes", responseBytes).Msg("could not do httputil.DumpResponse(res, true)")

		return nil, errWrapped
	}

	// safe to reassign a new logger with the responseBytes
	logger = logger.With().Interface("responseBytes", responseBytes).Logger()

	logger.Trace().Msg("successfully completed DebugTransport.RoundTrip")

	return res, err
}

// Logger - Returns a contextual logger.
func (d *DebugTransport) Logger() *zerolog.Logger {
	logger := d.LoggerParent.
		With().Str("type", "DebugTransport").Logger().
		With().Interface("DebugTransport.http.Transport", d.Transport).Logger()

	return &logger
}
