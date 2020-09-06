package everyday

import (
	. "leetcode/src/leetcode/tree"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {

	result := &[]string{}
	bytes := []int{}
	binaryTreePathsRecur(root, result, bytes)
	return *result
}
func binaryTreePathsRecur(root *TreeNode, s *[]string, b []int) {
	if root == nil {
		return
	}
	b = append(b, root.Val)
	if root.Left == nil && root.Right == nil {
		strs := []string{}
		for i := 0; i < len((b)); i++ {
			strs = append(strs, strconv.Itoa((b)[i]))
		}

		r := strings.Join(strs, "->")
		*s = append(*s, r)

	}
	binaryTreePathsRecur(root.Left, s, b)
	binaryTreePathsRecur(root.Right, s, b)
	return

}

// 60. 第k个排列
func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	row := k / factorial(n-1)
	col := k % factorial(n-1)
	return string(row+1) + getPermutation(n-1, col)
}

func factorial(n int) int {
	var facVal = 1
	if n < 0 {

	} else {
		for i := 1; i <= n; i++ {
			facVal *= i
		}
	}
	return facVal

}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}
