package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// 数组也是值传递 ,copy一分，，数组整个copy，不是仅仅拷贝元素地址
func TestT1(t *testing.T) {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		require.Equal(t, arr[0], 7)
	}(x)

	require.Equal(t, x[0], 1)
}

// slie的底层为数组，单slice放的指向数组的指针，，传递依然是值传递，拷贝，，只不过拷贝的是指针
func TestT2(t *testing.T) {
	x := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 7
		require.Equal(t, arr[0], 7)
	}(x)
	require.Equal(t, x[0], 7)
}

func TestT3(t *testing.T) {
	x := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 7           // 改变x
		arr = append(arr, 8) // append后，，slice指向的新的数组了，so，改变
		arr[1] = 9
		require.Equal(t, arr[0], 7)
		require.Equal(t, arr[1], 9)
	}(x)
	require.Equal(t, x[0], 7) //——|
	require.Equal(t, x[1], 2)
}
