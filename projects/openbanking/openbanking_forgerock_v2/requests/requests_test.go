package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Requests(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal("banaio-openbankingforgerock/3.0 (https://github.com/banaio/openbankingforgerock)", userAgent)
}
