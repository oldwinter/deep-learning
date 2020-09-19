/*
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_removeNthFromEnd(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	f := MakeLinkedList(in)
	result := removeNthFromEnd(f, 2)
	t.Log(PrintLinkedList(result))
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := &ListNode{}
	p1 := h
	p2 := h
	for i := 0; i < n; i++ {
		p2.Next = head
		head = head.Next
		p2 = p2.Next
	}
	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p1.Next = p1.Next.Next

	return h.Next
}
