/*
给定一个带有头结点 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。



示例 1：

输入：[1,2,3,4,5]
输出：此列表中的结点 3 (序列化形式：[3,4,5])
返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
注意，我们返回了一个 ListNode 类型的对象 ans，这样：
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.
示例 2：

输入：[1,2,3,4,5,6]
输出：此列表中的结点 4 (序列化形式：[4,5,6])
由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。


提示：

给定链表的结点数介于 1 和 100 之间。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/middle-of-the-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deep-learning/leetcode/common"
)

func Test_middleNode(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	f := MakeLinkedList(in)
	result := middleNode(f)
	t.Log(PrintLinkedList(result))

	in2 := []int{1, 2, 3, 4, 5, 6}
	f2 := MakeLinkedList(in2)
	result2 := middleNode(f2)
	t.Log(PrintLinkedList(result2))
}

func middleNode(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	for p2.Next != nil {
		p1 = p1.Next
		if p2.Next.Next == nil {
			break
		}
		p2 = p2.Next.Next
	}
	return p1
}
