package algorithm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		result float64
	}{
		{nil, []int{1}, 1},
		{nil, []int{1, 2}, 1.5},
		{nil, []int{1, 2, 3, 4}, 2.5},
		{[]int{2}, []int{1, 3}, 2},
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 3}, []int{2, 4}, 2.5},
		{[]int{1, 2, 3}, []int{1, 2}, 2},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}

	for i, tt := range tests {
		actual := FindMedianSortedArrays(tt.nums1, tt.nums2)
		require.Equalf(t, actual, tt.result, "The %d-th case: FindMedianSortedArrays(%v, %v), got %f, expected %f",
			i+1, tt.nums1, tt.nums2, actual, tt.result)
	}
}
