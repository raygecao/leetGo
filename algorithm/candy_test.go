package algorithm

import (
	"fmt"
	"testing"
)

func TestCandy(t *testing.T) {
	a := []int{1, 0, 2}
	b := []int{1, 2, 2}
	fmt.Println(Candy(a), Candy(b))
}
