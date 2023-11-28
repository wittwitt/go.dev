package ifelse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// defer retrun 执行顺序和作用范围

func test1() int { // return i, 将返回值写入返回堆栈，i已经不作用了
	i := 1
	defer func() {
		i++
	}()
	return i
}

func test2() int { //
	i := 1
	defer func(i int) {
		i++
	}(i)
	return i
}

func test3() (i int) { // 堆栈变量i
	i = 1
	defer func() {
		i++
	}()
	return i
}

func Test_defer1(t *testing.T) {
	require.Equal(t, test1(), 1)
	require.Equal(t, test2(), 1)
	require.Equal(t, test3(), 2)
}
