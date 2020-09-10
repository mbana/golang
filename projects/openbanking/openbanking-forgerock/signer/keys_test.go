package signer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/banaio/openbankingforgerock/config"
)

func Test_NewKeys(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	config := &config.Config{}
	keys, err := NewKeys(config)
	assert.Error(err)
	assert.Nil(keys)
}
