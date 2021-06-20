package tree

import (
	"fmt"

	"leetGo/linknode"
)

//
//给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
//
//
//
//示例：
//
//输入：[1,2,3,4,5,null,7,8]
//
//1
///  \
//2    3
/// \    \
//4   5    7
///
//8
//
//输出：[[1],[2,3],[4,5,7],[8]]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/list-of-depth-lcci
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func ListOfDepth(tree *TreeNode) []*linknode.ListNode {
	if tree == nil {
		return nil
	}
	queue := []*TreeNode{tree}
	var lists []*linknode.ListNode
	length := len(queue)
	var list, node *linknode.ListNode
	for len(queue) != 0 {
		root := queue[0]
		queue = queue[1:]
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		length--
		if list == nil {
			list = &linknode.ListNode{Val: root.Val}
			node = list
		} else {
			list.Next = &linknode.ListNode{Val: root.Val}
			list = list.Next
		}
		if length == 0 {
			lists = append(lists, node)
			list = nil
			length = len(queue)
			fmt.Println(length)
		}
	}
	return lists
}
