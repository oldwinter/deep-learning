/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。



示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。



注意：两结点之间的路径长度是以它们之间边的数目表示。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_diameterOfBinaryTree(t *testing.T) {

}
func diameterOfBinaryTree(root *TreeNode) int {
	var maxD = 0
	diameterOfBinaryTreeDFS(root, &maxD)
	return maxD
}
func diameterOfBinaryTreeDFS(root *TreeNode, maxD *int) int {
	if root == nil {
		return 0
	}

	leftLength := diameterOfBinaryTreeDFS(root.Left, maxD)
	rightLength := diameterOfBinaryTreeDFS(root.Right, maxD)
	if leftLength+rightLength > *maxD {
		*maxD = leftLength + rightLength
	}

	if leftLength > rightLength {
		return leftLength + 1
	} else {
		return rightLength + 1
	}

}
