package linkedlist

import "strings"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeLinkedList(sources []int) *ListNode {

	var head *ListNode = &ListNode{}
	temp := head
	for _, source := range sources {
		t := &ListNode{source, nil}
		temp.Next = t
		temp = temp.Next
	}

	return head.Next
}
func PrintLinkedList(head *ListNode) []int {
	a := []int{}
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}
	return a
}

// 206 反转链表 迭代
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

//141 环形链表
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}
	//哈希表存历史数据
	set := make(map[*ListNode]bool)

	for head != nil && !set[head] {

		set[head] = true
		head = head.Next
	}
	if head == nil {
		return false
	} else {
		return true
	}
}

// 21 合并两个有序链表
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

//19. 删除链表的倒数第N个节点
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

// 876 链表的中间节点
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

// 2. 两数相加
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

// 3. 无重复字符的最长子串
// func lengthOfLongestSubstring(s string) int {
// 	// 记录第i位起始的最长子串
// 	temp := map[int]int{}
// 	// 记录当前子串已经已经出现过的字符
// 	str := map[byte]bool{}
// 	for i := 0; i < len(s); i++ {
// 		j := i
// 		for ; j < len(s) && str[s[j]] == false; j++ {

// 			str[s[j]] = true

// 		}
// 		temp[i] = len(str)
// 		str = map[byte]bool{}
// 	}
// 	max := 0
// 	for _, v := range temp {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	return max
// }
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[left:i], string(s[i]))
		if index == -1 {
			if right < i+1 {
				right++
			}

		} else {
			left += index + 1
			right += index + 1

		}
	}
	return right - left
}
