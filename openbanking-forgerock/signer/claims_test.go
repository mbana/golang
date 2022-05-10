package signer

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func Test_Claims(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal([]string{
		"authorization_code",
		"refresh_token",
		"client_credentials",
	}, DefaultGrantTypes)
}
func Test_NewClaims(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	claims := NewClaims("AAC_a6550689-fd7f-438b-b10e-a37fd0804676")
	claimsJSON, err := json.Marshal(claims)

	assert.NoError(err)
	assert.JSONEq(`
{
    "id_token": {
        "openbanking_intent_id": {
            "essential": true,
            "value": "AAC_a6550689-fd7f-438b-b10e-a37fd0804676"
        },
        "acr": {
            "essential": true,
            "value": "urn:openbanking:psd2:sca"
        }
    },
    "user_info": {
        "openbanking_intent_id": {
            "essential": true,
            "value": "AAC_a6550689-fd7f-438b-b10e-a37fd0804676"
        }
    }
}
	`, string(claimsJSON))
}
