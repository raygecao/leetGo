package linknode

//编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。
//
//示例:
//
//输入: head = 3->5->8->5->10->2->1, x = 5
//输出: 3->1->2->10->5->5->8
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/partition-list-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func Partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	node, pre := head, head
	for {
		pre, node = node, node.Next
		if node == nil {
			break
		}
		if node.Val < x {
			pre.Next = node.Next
			node.Next, head = head, node
			node = pre
		}
	}
	return head
}
