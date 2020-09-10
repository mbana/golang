package pcsa

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPCSA(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	require.NotNil(NewPCSA(16))
}

func TestPCSALeadingOnes32(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	require.Equal(30, LeadingOnes32(math.MaxUint32-3))
	require.Equal(30, LeadingOnes32(math.MaxUint32-2))
	require.Equal(31, LeadingOnes32(math.MaxUint32-1))
	require.Equal(32, LeadingOnes32(math.MaxUint32))
}
