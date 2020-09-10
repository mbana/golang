package oidcdiscovery

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OpenIDConfiguration(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	filename := "testdata/well-known_openid-configuration.json"
	file, err := ioutil.ReadFile(filename)
	assert.NoError(err)

	config := &OpenIDConfiguration{}
	assert.NoError(json.Unmarshal(file, config))

	configJSON, err := json.MarshalIndent(config, "", "  ")
	assert.NoError(err)
	assert.JSONEq(string(file), string(configJSON))
}
