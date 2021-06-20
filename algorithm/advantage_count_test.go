package algorithm

import (
	"fmt"
	"testing"
)

func TestAdvantageCount(t *testing.T) {
	nums := []int{5, 1, 6, 7, 7, 2, 4}
	indexes := []int{0, 1, 2, 3, 4, 5, 6}
	fastSort(nums, indexes, 0, 6)
	fmt.Println(nums)
	fmt.Println(indexes)

	A := []int{5621, 1743, 5532, 3549, 9581}
	B := []int{913, 9787, 4121, 5039, 1481}
	fmt.Println(AdvantageCount(A, B))
	A = []int{12, 24, 8, 32}
	B = []int{13, 25, 32, 11}
	fmt.Println(AdvantageCount(A, B))

}
