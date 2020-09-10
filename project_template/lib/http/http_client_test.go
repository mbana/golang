package lib

import (
	"net/http"
	"testing"

	"github.com/banaio/golang/project_template/lib/utils"
)

func Test_HTTPClient_Template(t *testing.T) {
	assert, require, logger := utils.MakeTestCaseDependencies(t)
	_, _, _ = assert, require, logger
	t.Skipf("test=%+v is intentionally ignored it serves as a template to copy when making another test", t.Name())

	require.Equal("true", "true")
	assert.Equal("true", "true")
}

func Test_HTTPClient_SetDefaultHeaders(t *testing.T) {
	assert, require, logger := utils.MakeTestCaseDependencies(t)
	_, _, _ = assert, require, logger

	url := TestURL
	body := NewBodyEmpty()
	req, err := http.NewRequest(http.MethodGet, url, body)
	assert.NoError(err)
	assert.NotNil(req)

	SetDefaultHeaders(req)
	assert.NotNil(req)
	assert.NotNil(req.Header)

	expected := http.Header{
		HeaderAcceptKey:       {HeaderAcceptValueAll},
		HeaderCacheControlKey: {HeaderCacheControlValueNoCache},
		HeaderUserAgentKey:    {HeaderUserAgentValue},
	}
	actual := req.Header

	assert.Equal(expected, actual)
}
