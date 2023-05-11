// Package nil nil
package nil

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

// Da da
type Da struct {
}

func Test_t1(t *testing.T) {

	var some1 interface{}
	t1 := ToIS(some1)
	require.Equal(t, true, t1.pt == 0)
	require.Equal(t, true, t1.pv == 0)

	some2 := 1
	t2 := ToIS(some2)

	some22 := 12
	t22 := ToIS(some22)

	require.Equal(t, true, t2.pt == t22.pt)
	require.Equal(t, true, t2.pv != t22.pv)

	var some3 *Da
	t3 := ToIS(some3)

	var some33 *Da
	t33 := ToIS(some33)

	require.Equal(t, true, t3.pt == t33.pt)
	require.Equal(t, true, t3.pv == 0)
	require.Equal(t, true, t33.pv == 0)
}

// IS IS
type IS struct {
	pt uintptr
	pv uintptr
}

// ToIS ToIS
func ToIS(i interface{}) IS {
	return *(*IS)(unsafe.Pointer(&i))
}
