package linkedlist

//练习题LeetCode对应编号：206，141，21，19，876。大家可以去练习，另外建议作者兄每章直接给出LC的题目编号或链接方便大家练习。

import (
	"testing"
)

func TestLinkedlist(t *testing.T) {
	if true {
		t.Log(lengthOfLongestSubstring("pwwkew"))

	}

	if false {
		in := []int{2, 4, 3}
		in2 := []int{2, 6}
		f1 := MakeLinkedList(in)
		f2 := MakeLinkedList(in2)
		f := addTwoNumbers(f1, f2)
		t.Log(PrintLinkedList(f))
		// t.Log(f)
	}

	if false {
		// 206 反转链表 迭代
		var f5 *ListNode = &ListNode{5, nil}
		var f4 *ListNode = &ListNode{4, f5}
		var f3 *ListNode = &ListNode{3, f4}
		var f2 *ListNode = &ListNode{2, f3}
		var f1 *ListNode = &ListNode{1, f2}
		var result *ListNode = reverseList(f1)
		t.Log(result.Val, result.Next.Val, result.Next.Next.Val)

		// 141 环形链表

		var f5_141 *ListNode = &ListNode{-4, nil}
		var f4_141 *ListNode = &ListNode{0, nil}
		var f3_141 *ListNode = &ListNode{2, nil}
		var f2_141 *ListNode = &ListNode{3, nil}

		f5_141.Next = f3_141
		f4_141.Next = f5_141
		f3_141.Next = f4_141
		f2_141.Next = f3_141

		result_141 := hasCycle(f2_141)
		t.Log(result_141)

		var f5_141_1 *ListNode = &ListNode{1, nil}
		var f4_141_1 *ListNode = &ListNode{2, nil}
		f5_141_1.Next = f4_141_1

		result_141_1 := hasCycle(f5_141_1)
		t.Log(result_141_1)
	}

	if false {
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

	if false {
		// 19 删除链表的倒数第N个节点
		in := []int{1, 2, 3, 4, 5}
		f := MakeLinkedList(in)
		result := removeNthFromEnd(f, 2)
		t.Log(PrintLinkedList(result))
	}

	if false {
		// 19 删除链表的倒数第N个节点
		in := []int{1, 2, 3, 4, 5}
		f := MakeLinkedList(in)
		result := middleNode(f)
		t.Log(PrintLinkedList(result))

		in2 := []int{1, 2, 3, 4, 5, 6}
		f2 := MakeLinkedList(in2)
		result2 := middleNode(f2)
		t.Log(PrintLinkedList(result2))

	}

}
