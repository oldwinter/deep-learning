/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deep-learning/leetcode/common"
)

func Test_addTwoNumbers(t *testing.T) {
	in := []int{2, 4, 3}
	in2 := []int{2, 6}
	f1 := MakeLinkedList(in)
	f2 := MakeLinkedList(in2)
	f := addTwoNumbers(f1, f2)
	t.Log(PrintLinkedList(f))
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	bitFlag := 0
	result := &ListNode{}
	head := result
	for l1 != nil && l2 != nil {
		temp := l1.Val + l2.Val + bitFlag
		if temp >= 10 {
			node := &ListNode{temp % 10, nil}
			bitFlag = 1
			result.Next = node
		} else {
			// bitFlag =0
			node := &ListNode{temp, nil}
			bitFlag = 0
			result.Next = node
		}
		result = result.Next
		l1 = l1.Next
		l2 = l2.Next

	}
	for l1 != nil {
		temp := l1.Val + bitFlag
		if temp >= 10 {
			node := &ListNode{temp % 10, nil}
			bitFlag = 1
			result.Next = node
		} else {
			// bitFlag =0
			node := &ListNode{temp, nil}
			bitFlag = 0
			result.Next = node
		}
		result = result.Next
		l1 = l1.Next
	}
	for l2 != nil {
		temp := l2.Val + bitFlag
		if temp >= 10 {
			node := &ListNode{temp % 10, nil}
			bitFlag = 1
			result.Next = node
		} else {
			// bitFlag =0
			node := &ListNode{temp, nil}
			bitFlag = 0
			result.Next = node
		}
		result = result.Next
		l2 = l2.Next
	}
	if bitFlag == 1 {
		node := &ListNode{1, nil}
		// bitFlag = 0
		result.Next = node
	}
	return head.Next

}
