package signer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewSigner(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	keys, err := NewSigner(nil, nil)
	assert.Error(err)
	assert.Nil(keys)
}
