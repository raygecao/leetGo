package linknode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1     *ListNode
		l2     *ListNode
		result *ListNode
	}{
		{nil, makeLinkNode(1, 2), makeLinkNode(1, 2)},
		{makeLinkNode(1, 2), makeLinkNode(0), makeLinkNode(1, 2)},
		{makeLinkNode(5), makeLinkNode(5), makeLinkNode(0, 1)},
		{makeLinkNode(1), makeLinkNode(9, 9), makeLinkNode(0, 0, 1)},
		{makeLinkNode(2, 4, 3), makeLinkNode(5, 6, 4), makeLinkNode(7, 0, 8)},
		{makeLinkNode(2, 4, 3), makeLinkNode(5, 6, 6), makeLinkNode(7, 0, 0, 1)},
		{makeLinkNode(2, 4, 3, 4, 6), makeLinkNode(5, 6, 6), makeLinkNode(7, 0, 0, 5, 6)},
		{makeLinkNode(2, 4, 3, 9), makeLinkNode(5, 6, 6), makeLinkNode(7, 0, 0, 0, 1)},
	}

	for i, tt := range tests {
		actual := getListNums(AddTwoNumbers(tt.l1, tt.l2))
		expected := getListNums(tt.result)
		require.ElementsMatch(t, actual, expected, "the %d-th cases : AddTwoNumbers(%v, %v), got %v, expect %v",
			i+1, getListNums(tt.l1), getListNums(tt.l2), actual, expected)
		actualEnhanced := getListNums(AddTwoNumbersEnhanced(tt.l1, tt.l2))
		require.ElementsMatch(t, actualEnhanced, expected, "the %d-th cases : AddTwoNumbersEnhanced(%v, %v), got %v, expect %v",
			i+1, getListNums(tt.l1), getListNums(tt.l2), actualEnhanced, expected)
	}
}
