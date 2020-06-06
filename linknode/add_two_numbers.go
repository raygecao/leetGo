package linknode

// AddTwoNumbers add two link list and get a summed link list.
// Each node only contains a number.
// The number is organized in desc order.
// The number won't be start with 0, except it's exact zero.
// e.g. AddTwoNumbers((2->4->3), (5->6->4)) = (7->0->8)
// ref link: https://leetcode-cn.com/problems/add-two-numbers
// leecode-2
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// if one list is nil, return the other
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	cur1, cur2 := l1, l2
	head := new(ListNode)
	cursor := head
	var step int

	// both lists are not nil
	for {
		tmp := new(ListNode)
		if sum := cur1.Val + cur2.Val + step; sum >= 10 {
			tmp.Val = sum - 10
			step = 1
		} else {
			tmp.Val = sum
			step = 0
		}
		cursor.Next = tmp
		cursor = cursor.Next

		cur1, cur2 = cur1.Next, cur2.Next

		// two list both are ended
		if cur1 == nil && cur2 == nil {
			if step != 0 {
				cursor.Next = &ListNode{Val: step}
				return head.Next
			}
			return head.Next
		}

		// add a zero node to the shorter node
		if cur1 == nil {
			cur1 = &ListNode{Val: 0}
		}
		if cur2 == nil {
			cur2 = &ListNode{Val: 0}
		}
	}
}

// AddTwoNumbersEnhanced solves AddTwoNumbers in a way saving memory overhead
// which destroys origin lists
func AddTwoNumbersEnhanced(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	cur1, cur2 := l1, l2
	var step int

	// tmp is the front node of cur1
	var tmp *ListNode
	for cur1 != nil && cur2 != nil {
		if sum := cur1.Val + cur2.Val + step; sum >= 10 {
			cur1.Val = sum - 10
			step = 1
		} else {
			cur1.Val = sum
			step = 0
		}
		tmp = cur1
		cur1, cur2 = cur1.Next, cur2.Next
	}

	// len(l1) == len(l2)
	if cur1 == nil && cur2 == nil {
		if step > 0 {
			tmp.Next = &ListNode{Val: step}
		}
		return l1
	}

	// is cur1 is shorter than cur2, move cur2 tail list to cur1
	if cur1 == nil {
		cur1 = cur2
		tmp.Next = cur1
	}

	for cur1 != nil {
		if step+cur1.Val < 10 {
			cur1.Val += step
			return l1
		}
		cur1.Val = 0
		step = 1
		tmp = cur1
		cur1 = cur1.Next
	}
	if step > 0 {
		tmp.Next = &ListNode{Val: step}
	}
	return l1
}
