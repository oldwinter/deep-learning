package everyday

import (
	. "leetcode/src/leetcode/tree"
	"testing"
)

func TestEveryday(t *testing.T) {

	if true {
		in := []int{1, 1, 1, 2, 2, 3}
		t.Log(topKFrequent(in, 2))

	}
	if false {
		preorder := []int{3, 9, 20, 15, 7}
		inorder := []int{9, 3, 15, 20, 7}
		result := buildTree(preorder, inorder)
		t.Log(binaryTreePaths(result))
	}
}

// 剑指 Offer 07. 重建二叉树 假设不含重复数字
func buildTree(preorder []int, inorder []int) *TreeNode {
	// result := TreeNode{}
	// for  {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0], Left: nil, Right: nil}

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

	leftnode := buildTree(preorder[1:mid], inorder[:mid-1])
	rightnode := buildTree(preorder[mid:], inorder[mid:])

	root.Left = leftnode
	root.Right = rightnode
	return root
	// }
}
