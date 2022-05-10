package requests

import (
	"io"
	"net/http"
)

const (
	userAgent = "banaio-openbankingforgerock/3.0 (https://github.com/banaio/openbankingforgerock)"
)

// MIME types
const (
	MIMEApplicationJSON = "application/json"
	MIMEApplicationJWT  = "application/jwt"
	MIMEApplicationForm = "application/x-www-form-urlencoded"
)

// Headers standard HTTP
const (
	HeaderAccept        = "Accept"
	HeaderAuthorization = "Authorization"
	HeaderCacheControl  = "Cache-Control"
	HeaderContentType   = "Content-Type"
	HeaderUserAgent     = "User-Agent"
)

// Headers FAPI
const (
	HeaderXFAPICustomerIPAddress      = "x-fapi-customer-ip-address"
	HeaderXFAPICustomerLastLoggedTime = "x-fapi-customer-last-logged-time"
	HeaderXFAPIFinancialID            = "x-fapi-financial-id"
	HeaderXFAPIInteractionID          = "x-fapi-interaction-id"
	HeaderXIdempotencyKey             = "x-idempotency-key"
	HeaderXJWSSignatureKey            = "x-jws-signature"
)

func NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	SetDefaultHeaders(req)
	return req, nil
}

func SetDefaultHeaders(req *http.Request) {
	req.Header.Set(HeaderCacheControl, "no-cache")
	req.Header.Set(HeaderUserAgent, userAgent)
}
