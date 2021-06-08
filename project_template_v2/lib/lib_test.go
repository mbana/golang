package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Lib_Assert(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(true, true, "they should be equal")
}

func Test_Lib_Require(t *testing.T) {
	require := require.New(t)

	require.Equal(true, true, "they should be equal")
}
