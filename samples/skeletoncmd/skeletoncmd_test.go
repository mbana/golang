package skeletoncmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkeletonCmd(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	require.Equal(true, SkeletonCmd())
}
