// Package lib - this parts contains the *http.Client used to make requests.
//
// Could possibly just use another HTTP library that has some of these constants defined ...
//
// Since the std. library doesn't contain these things, I need to redefine them.
// A good library to use would be https://github.com/labstack/echo/blob/master/echo.go.
// In particular see these lines, https://github.com/labstack/echo/blob/master/echo.go#L153.
package lib

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/banaio/golang/project_template/lib/logger"
	"github.com/banaio/golang/project_template/lib/utils"
)

const (
	// ContentTypeHTMLUTF8 - Content-Type: Possible returns types.
	//
	// Examples:
	//
	// * "content-type: text/html; charset=UTF-8"
	// * "content-type: text/html; charset=utf-8"
	// * "content-type: text/html;charset=utf-8"
	//
	// All the examples above are valid. Note the spacing and casing does not seem to matter.
	// ContentTypeHTMLUTF8 = "content-type: text/html; charset=UTF-8"
	// Remove above line.
	ContentTypeHTMLUTF8 = "text/html; charset=UTF-8"
)

// ContentTypeHTMLValidValues -  Content-Type: Possible returns types.
// There are possibly more - I need to consult the RFCs.
// `ContentTypeHTMLValidValues` is a global variable (gochecknoglobals)
// nolint:gochecknoglobals
var ContentTypeHTMLValidValues = []string{
	// // Remove first block with `content-type`.
	// "content-type: text/html; charset=UTF-8",
	// "content-type: text/html; charset=utf-8",
	// "content-type: text/html;charset=utf-8",
	"text/html; charset=UTF-8",
	"text/html; charset=utf-8",
	"text/html;charset=utf-8",
	"text/html; charset=UTF-8",
}

// HTTP standard headers.
const (
	HeaderAcceptKey       = "Accept"
	HeaderCacheControlKey = "Cache-Control"
	HeaderUserAgentKey    = "User-Agent"
	HeaderContentTypeKey  = "Content-Type"
)

// Header standard headers values that we use.
const (
	HeaderAcceptValueAll           = "*/*"
	HeaderCacheControlValueNoCache = "no-cache"
)

const (
	Version = "0.0.1"
	Website = "https://bana.io"
)

// HeaderUserAgentValue - `User-Agent` to use. Panic if we fail to get it.
// Example: "github.com/banaio/golang/project_template/0.0.1 (https://bana.io)".
//
// `HeaderUserAgentValue` is a global variable (gochecknoglobals)go-lint
// nolint:gochecknoglobals
var HeaderUserAgentValue = GetHeaderUserAgentValue()

// GetHeaderUserAgentValueModuleName - generally a bad idea to exit in the function.
// However, if we can't determine module name, something is really wrong..
// Returns something like "github.com/banaio/golang/project_template/0.0.1 (https://bana.io)".
func GetHeaderUserAgentValue() string {
	verbose := make([]bool, zerolog.WarnLevel)
	logWithTimestamps := false
	logger := logger.NewLogger(verbose, logWithTimestamps)
	appName, err := utils.GetAppName(logger)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"%s%v%s",
			utils.ColourRed,
			errors.Wrapf(err, "lib.http_client.GetHeaderUserAgentValueModuleName could not get appName"),
			utils.ColourReset,
		)
		os.Exit(utils.ExitCodeErr1)
	}

	return fmt.Sprintf("%s/%s (%s)", appName, Version, Website)
}

func SetDefaultHeaders(req *http.Request) *http.Request {
	// Could just do:
	// req.Header = http.Header{
	// 	HeaderAcceptKey:       {HeaderAcceptValueAll},
	// 	HeaderCacheControlKey: {HeaderCacheControlValueNoCache},
	// 	HeaderUserAgentKey:    {HeaderUserAgentValue},
	// }
	// But it's safer to do it as below.
	//
	// Should possibly add???
	// 	req.Header.Set("Connection", "keep-alive")
	req.Header.Set(HeaderAcceptKey, HeaderAcceptValueAll)
	req.Header.Set(HeaderCacheControlKey, HeaderCacheControlValueNoCache)
	req.Header.Set(HeaderUserAgentKey, HeaderUserAgentValue)

	return req
}
