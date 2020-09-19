package common

// ListNode represents a node in a linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// MakeLinkedList 用数组构造链表
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

// PrintLinkedList 按数组形式打印链表
func PrintLinkedList(head *ListNode) []int {
	results := []int{}
	for head != nil {
		results = append(results, head.Val)
		head = head.Next
	}
	return results
}

// NewListNode returns a new list node.
func NewListNode(v int) *ListNode {
	return &ListNode{v, nil}
}

// AddNext adds a next node to the end of list.
func (l *ListNode) AddNext(v int) {
	for n := l; n != nil; n = n.Next {
		if n.Next == nil {
			new := NewListNode(v)
			n.Next = new
			break
		}
	}
}

// LinkedListToSlice converts a linked list into an array of integer.
func LinkedListToSlice(node *ListNode) []int {
	out := []int{}

	for n := node; n != nil; n = n.Next {
		out = append(out, n.Val)
	}

	return out
}
