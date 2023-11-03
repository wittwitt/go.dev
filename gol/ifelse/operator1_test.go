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

	require.Equal(t, 0, 0%100)
	require.Equal(t, 1, 1%100)
	require.Equal(t, 50, 50%100)
	require.Equal(t, 0, 100%100)
	require.Equal(t, 1, 101%100)
}

func Test_yi(t *testing.T) {

	t.Log(uint64(512 << 10))
}
