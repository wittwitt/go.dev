package nil1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nil_relect(t *testing.T) {
	// c := *(*int)(nil)
	// t.Log(c)

	d := (*int)(nil)
	require.Equal(t, true, d == nil)
}
