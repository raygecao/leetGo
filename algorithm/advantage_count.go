package algorithm

import (
	"sort"
)

// 4, 2, 10
// 1, 3, 5, 7, 9  => 5, 1, 3, 7, 9 => 5, 3, 1, 7, 9
func AdvantageCount(A []int, B []int) []int {
	if len(A) < 2 {
		return A
	}
	sort.Ints(A)
	indexes := make([]int, len(B))
	for i := range B {
		indexes[i] = i
	}
	last := len(B) - 1
	fastSort(B, indexes, 0, last)
	for i := 0; i <= last; {
		if A[i] <= B[i] {
			min := A[0]
			A = append(A[1:], min)
			last--
		} else {
			i++
		}
	}
	fastSort(indexes, A, 0, len(B)-1)
	return A
}

func fastSort(nums, indexes []int, start, end int) {
	if end <= start {
		return
	}
	target, index := nums[start], indexes[start]
	i, j := start, end
	for i < j {
		for i < j && nums[j] >= target {
			j--
		}
		nums[i] = nums[j]
		indexes[i] = indexes[j]
		for i < j && nums[i] <= target {
			i++
		}
		nums[j] = nums[i]
		indexes[j] = indexes[i]
	}
	nums[i] = target
	indexes[i] = index
	fastSort(nums, indexes, start, i-1)
	fastSort(nums, indexes, i+1, end)
}
