package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxSumBST gets the max sum of sub BST
func MaxSumBST(root *TreeNode) int {
	sum := new(int)
	BSTStatus(root, sum)
	return *sum
}

func BSTStatus(root *TreeNode, m *int) (isBST bool, min, max, sum int) {
	if root == nil {
		return true, 40001, -40001, 0
	}

	isLeft, minLeft, maxLeft, sumLeft := BSTStatus(root.Left, m)
	isRight, minRight, maxRight, sumRight := BSTStatus(root.Right, m)
	// not BST
	if !isLeft || !isRight || maxLeft >= root.Val || minRight <= root.Val {
		return false, 0, 0, 0
	}
	isBST = true
	// get min and max of the root
	min, max = minLeft, maxRight
	if root.Val < min {
		min = root.Val
	}
	if root.Val > max {
		max = root.Val
	}
	// get the sum of the root
	sum = root.Val + sumLeft + sumRight

	if sum > *m {
		*m = sum
	}

	return
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	if mid+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}
	return root
}
