package tree

import (
	"fmt"
	"testing"
)

func TestMaxSumBST(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 4}
	root.Left = left
	p := MaxSumBST(root)
	fmt.Println(p)
}
