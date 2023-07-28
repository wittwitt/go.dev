package struct1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type InitCat struct {
	array1 [100]int
	slice1 []int
}

func Test_init(t *testing.T) {
	ic := InitCat{}
	require.Equal(t, 100, len(ic.array1))
	require.Equal(t, 0, len(ic.slice1))
}
