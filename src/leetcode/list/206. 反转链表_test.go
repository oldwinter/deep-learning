/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_reverseList(t *testing.T) {
	var f5 *ListNode = &ListNode{5, nil}
	var f4 *ListNode = &ListNode{4, f5}
	var f3 *ListNode = &ListNode{3, f4}
	var f2 *ListNode = &ListNode{2, f3}
	var f1 *ListNode = &ListNode{1, f2}
	var result *ListNode = reverseList(f1)
	t.Log(result.Val, result.Next.Val, result.Next.Next.Val)
}
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode = nil
	for head != nil {
		p := head.Next
		head.Next = prev
		prev = head
		head = p

	}
	return prev
}
