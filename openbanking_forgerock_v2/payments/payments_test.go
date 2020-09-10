package payments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Payments(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	payments := NewPayments()
	assert.NotNil(payments)
}
