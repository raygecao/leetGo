package interview

import (
	"reflect"
	"testing"
)

func TestCopySlice(t *testing.T) {

	s1 := make([]int, 3)
	s1[0], s1[1], s1[2] = 1, 2, 3
	var s2, s3 []int
	s2 = append(s2, 1, 2, 3)
	s3 = append(s3, 1, 2, 3, 4)

	s := []struct {
		input   []int
		output1 []int
		output2 []int
	}{
		{nil, nil, nil},
		{
			[]int{1}, []int{1, 1}, []int{1, 2},
		},
		{
			[]int{1, 2}, []int{1, 2, 1}, []int{1, 2, 2},
		},
		{
			s1, []int{1, 2, 3, 1}, []int{1, 2, 3, 2},
		},
		{
			s2, []int{1, 2, 3, 2}, []int{1, 2, 3, 2},
		},
		{
			s3, []int{1, 2, 3, 4, 1}, []int{1, 2, 3, 4, 2},
		},
	}

	for _, tt := range s {
		if o1, o2 := CopySlice(tt.input); !reflect.DeepEqual(o1, tt.output1) || !reflect.DeepEqual(o2, tt.output2) {
			t.Errorf("CopySlice(%v), got(%v, %v), expect(%v, %v)", tt.input, o1, o2, tt.output1, tt.output2)
		}
	}
}
