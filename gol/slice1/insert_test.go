package slice1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_insert(t *testing.T) {

	// slice 插入操作

	slice := []int{1, 2, 3, 4, 5}

	{
		newSlice := []int{} // 正确做法是，现生成一个slice，然后，前+new_item+后
		newSlice = append(newSlice, slice[0:2]...)
		newSlice = append(newSlice, 6) // insert new item
		newSlice = append(newSlice, slice[2:]...)
		require.Equal(t, []int{1, 2, 6, 3, 4, 5}, newSlice)
	}

	{
		// 错误做法
		newSlice := append(slice[0:2], 6) // insert new item
		newSlice = append(newSlice, slice[2:]...)

		// 这里可能正确，可能错误
		// 当 append(slice[0:2], 6)， 修改的是slice指向的数组，，数组容量够的情况下，newslice和slice指向同一个数组，就不正确了
		// 不够的情况下，，新建，，就正确了

		require.NotEqual(t, []int{1, 2, 6, 3, 4, 5}, newSlice) // []int{1, 2, 6, 6, 4, 5}
	}
}

func insertSliceItem(slice []int, item int, index int) []int {
	// Ensure the index is within the bounds of the slice.
	if index < 0 || index > len(slice) {
		fmt.Println("Index out of range")
		return slice
	}

	// Create a new slice to accommodate the inserted item.
	result := append(slice[:index], append([]int{item}, slice[index:]...)...)

	return result
}

func Test_main(t *testing.T) {
	// Create a slice of integers

	slice := make([]int, 20) // len = 20, cap = 20
	slice = slice[:0]
	slice = append(slice, []int{1, 2, 3, 4, 5}...)

	// Insert the item 99 at index 2
	slice2 := insertSliceItem(slice, 99, 2)

	fmt.Println(slice)

	fmt.Println(slice2)
}
