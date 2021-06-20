package dfs

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(Permute(a))
	fmt.Println(Permute2(a))
}
