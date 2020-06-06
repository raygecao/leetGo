package linknode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func makeLinkNode(nums ...int) *ListNode {
	head := new(ListNode)
	cursor := head
	for _, num := range nums {
		tmp := &ListNode{Val: num}
		cursor.Next = tmp
		cursor = cursor.Next
	}
	return head.Next
}

func getListNums(l *ListNode) []int {
	var result []int
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return result
}
