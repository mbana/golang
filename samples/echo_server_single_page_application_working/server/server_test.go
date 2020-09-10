package server

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"bytes"
	"io"

	"github.com/labstack/echo"
)

func TestServerApiHello(t *testing.T) {
	require := require.New(t)

	server := NewServer()
	defer func() {
		require.NoError(server.Shutdown(context.TODO()))
	}()
	require.NotNil(server)

	// make the request
	code, body, headers := request(
		http.MethodGet,
		"/api/hello?pretty",
		nil,
		server,
	)

	// do assertions
	require.Equal(http.StatusOK, code)
	require.Len(headers, 1)
	require.Equal("text/plain; charset=UTF-8", headers["Content-Type"][0])

	require.NotNil(body)
	require.Equal("Hello, World!", body.String())
}

func TestServerGetFavicon(t *testing.T) {
	require := require.New(t)

	server := NewServer()
	defer func() {
		require.NoError(server.Shutdown(context.TODO()))
	}()
	require.NotNil(server)

	// make the request
	code, body, headers := request(
		http.MethodGet,
		"/favicon.ico",
		nil,
		server,
	)

	// do assertions
	require.Equal(http.StatusOK, code)
	require.Len(headers, 4)
	t.Logf("headers=%#v=", headers)
	require.Contains(headers["Content-Type"][0], "icon")
	require.Equal([]string{"12"}, headers["Content-Length"])

	require.NotNil(body)
	require.Equal("favicon.ico\n", body.String())
}

func TestServerGetFaviconDoesNotExist(t *testing.T) {
	require := require.New(t)

	server := NewServer()
	defer func() {
		require.NoError(server.Shutdown(context.TODO()))
	}()
	require.NotNil(server)

	// make the request
	code, body, _ := request(
		http.MethodGet,
		"/favicon_does_not_exist.ico",
		nil,
		server,
	)

	// do assertions
	require.Equal(http.StatusNotFound, code)
	require.JSONEq(`{"message":"Not Found"}`, body.String())
}

func TestServerGetReturnsIndex(t *testing.T) {
	require := require.New(t)

	server := NewServer()
	defer func() {
		require.NoError(server.Shutdown(context.TODO()))
	}()
	require.NotNil(server)

	// make the request
	code, body, headers := request(
		http.MethodGet,
		"/",
		nil,
		server,
	)

	// do assertions
	require.Equal(http.StatusOK, code)
	require.Len(headers, 4)
	require.Equal([]string{"text/html; charset=utf-8"}, headers["Content-Type"])
	require.Equal([]string{"667"}, headers["Content-Length"])

	require.NotNil(body)

	expected, err := ioutil.ReadFile(filepath.Join("testdata", "index.html"))
	require.NoError(err)

	require.Equal(string(expected), body.String())
}

func request(method, path string, body io.Reader, server *echo.Echo) (int, *bytes.Buffer, http.Header) {
	req := httptest.NewRequest(method, path, body)
	rec := httptest.NewRecorder()

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	server.ServeHTTP(rec, req)

	return rec.Code, rec.Body, rec.Header()
}
