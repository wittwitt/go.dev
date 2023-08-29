package slice1

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sort1(t *testing.T) {

	{
		list := []int{3, 6, 1, 2, 9, 10, 0}

		sort.Slice(list, func(i, j int) bool {
			return list[i] < list[j]
		})
		require.Equal(t, []int{0, 1, 2, 3, 6, 9, 10}, list)
	}

	{
		list := []int{3, 6, 1, 2, 9, 10, 0}

		sort.Slice(list, func(i, j int) bool {
			return list[i] > list[j]
		})
		require.Equal(t, []int{10, 9, 6, 3, 2, 1, 0}, list)
	}
}
