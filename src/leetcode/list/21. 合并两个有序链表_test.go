/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。



示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_mergeTwoLists(t *testing.T) {
	// 206 反转链表 迭代
	var f1 *ListNode = &ListNode{1, nil}
	var f2 *ListNode = &ListNode{2, nil}
	var f3 *ListNode = &ListNode{4, nil}
	f1.Next = f2
	f2.Next = f3

	var g1 *ListNode = &ListNode{1, nil}
	var g2 *ListNode = &ListNode{3, nil}
	var g3 *ListNode = &ListNode{4, nil}
	g1.Next = g2
	g2.Next = g3
	result := mergeTwoLists(f1, g1)
	t.Log(PrintLinkedList(result))
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	result := head
	// if l1.Val < l2.Val {
	// 	head.Next = l1
	// } else {
	// 	head.Next = l2
	// }

	// head := If(.Val,l1,l2).(ListNode)
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {

			result.Next = l1
			l1 = l1.Next
		} else {
			result.Next = l2
			l2 = l2.Next
		}
		result = result.Next
	}

	for l1 != nil {
		result.Next = l1
		l1 = l1.Next
		result = result.Next
	}
	for l2 != nil {
		result.Next = l2
		l2 = l2.Next
		result = result.Next

	}

	return head.Next
}
