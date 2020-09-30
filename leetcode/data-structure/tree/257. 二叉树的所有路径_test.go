/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/oldwinter/deep-learning/leetcode/common"
)

func Test_binaryTreePaths(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	result := BuildTree(preorder, inorder)
	t.Log(binaryTreePaths(result))
}
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
