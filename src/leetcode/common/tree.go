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
