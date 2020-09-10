package server

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Single Page Application
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper: skipper,
		Root:    "testdata/",
		Index:   "index.html",
		HTML5:   true,
		Browse:  false,
	}))

	// Routes
	e.GET("/api/hello", hello)

	return e
}

// Skipper ensures that all requests not prefixed with `/api` get sent
// to the `middleware.Static` or `middleware.StaticWithConfig`.
// E.g., ensure that `/api` does not get handled by the
// the static middleware.
//
// Anything not prefix by `/api` will get get handled by
// `middleware.Static` or `middleware.StaticWithConfig`
func skipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/api")
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
