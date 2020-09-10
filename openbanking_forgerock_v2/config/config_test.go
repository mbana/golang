package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewConfig_Good(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	config, err := NewConfig("./testdata/config_good.yml")
	assert.NoError(err)
	assert.NotNil(config)
}

func Test_NewConfig_Bad(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	config, err := NewConfig("./testdata/config_bad.yml")
	assert.EqualError(err, "failed on config.Validate in NewConfig: filename=\"./testdata/config_bad.yml\": KID: cannot be blank; ORGANISATION_ID: cannot be blank; SSA: cannot be blank; SSID: cannot be blank; client_id: cannot be blank; encryption_private: cannot be blank; encryption_public: cannot be blank; register_response: cannot be blank; request_object_signing_alg: cannot be blank; signature_private: cannot be blank; signature_public: cannot be blank; transport_private: cannot be blank; transport_public: cannot be blank.")
	assert.Nil(config)
}
