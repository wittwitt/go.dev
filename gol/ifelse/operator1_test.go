package ifelse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_var(t *testing.T) {

	// :=
	// 相当于
	// var x type
	// x=some

	a := 100
	b := 200
	a, b, c, d := f1() // := , 最少有一个未声明，
	require.Equal(t, a, 11)
	require.Equal(t, b, 12)
	require.Equal(t, c, 13)
	require.Equal(t, d, true)

	{
		require.Equal(t, a, 11)

		a, x, _, _ := f2() //  :=范围限定在{}内， a就是新变量了，不会覆盖原有变量
		require.Equal(t, a, 1)
		require.Equal(t, x, 2)
		a = 123
	}

	require.Equal(t, a, 11)

}

func f1() (int, int, int, bool) {
	return 11, 12, 13, true
}
func f2() (int, int, int, bool) {
	return 1, 2, 3, true
}

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

	t.Log(4 << 20)
}

func Test_mod(t *testing.T) {

	i := 0x12
	j := 0x12

	t.Log(i ^ j)

}

func Test_you(t *testing.T) {
	var x int64 = 5
	var y int64 = 3

	t.Log(-x - y)
}
