package algorithm

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{1, 2, 3}, 5, []int{1, 2}},
		{[]int{1, 2, 3}, 4, []int{0, 2}},
		{[]int{1, 2, 3}, 10, nil},
		{nil, 5, nil},
	}

	for i, tt := range tests {
		if result := TwoSum(tt.nums, tt.target); !reflect.DeepEqual(result, tt.result) {
			t.Errorf("The %d-th case: TwoSum(%v, %d), got %v, expect %v", i+1, tt.nums, tt.target, result, tt.result)
		}
	}
}
