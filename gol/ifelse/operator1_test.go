package ifelse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestT1(t *testing.T) {
	require.Equal(t, 0, 1/5)
	require.Equal(t, 0, 3/5)
	require.Equal(t, 1, 5/5)
	require.Equal(t, 1, 6/5)
	require.Equal(t, 1, 8/5)
	require.Equal(t, 1, 9/5)
	require.Equal(t, 2, 10/5)
}
