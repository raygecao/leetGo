package stack

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	a := []int{7, 1, 5, 3, 6, 4}
	b := []int{7, 6, 4, 3, 1}
	fmt.Println(MaxProfit(a), MaxProfit(b))
}
