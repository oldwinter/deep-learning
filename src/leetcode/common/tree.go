package common

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode returns a new TreeNode.
func NewTreeNode(v int) *TreeNode {
	return &TreeNode{v, nil, nil}
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}

	if len(preorder) == 1 {
		return root
	}

	mid := 0
	for i, v := range inorder {
		if v == preorder[0] {
			mid = i
		}
	}
	mid++

	leftnode := BuildTree(preorder[1:mid], inorder[:mid-1])
	rightnode := BuildTree(preorder[mid:], inorder[mid:])

	root.Left = leftnode
	root.Right = rightnode
	return root
}

// 广度优先遍历，以数组形式存放
// 关键点：用一个队列，把树的层的前后给记住
func PrintTree(head *TreeNode) []int {
	if head == nil {
		return []int{}
	}
	result := []int{}
	a := []TreeNode{}
	a = append(a, *head)
	for len(a) != 0 {
		head := a[0]
		result = append(result, head.Val)
		a = a[1:]
		if head.Left != nil {
			a = append(a, *head.Left)
			// result = append(result, head.Left.Val)
		} else {
			// result = append(result, -1)
		}
		if head.Right != nil {
			a = append(a, *head.Right)
		} else {
			// result = append(result, -1)
		}

	}
	return result
}

func PrintTree2(root *TreeNode) [][]int {
	results := [][]int{}
	if root == nil {
		return results
	}
	levelQueue := []TreeNode{*root}
	level := 0
	for len(levelQueue) != 0 {
		a := []TreeNode{}
		results = append(results, []int{})
		for _, v := range levelQueue {
			// 取值
			results[level] = append(results[level], v.Val)
			// 入队
			if v.Left != nil {
				a = append(a, *v.Left)
			}
			if v.Right != nil {
				a = append(a, *v.Right)
			}

		}
		level++
		levelQueue = a

	}

	return results
}
