/*
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


限制：

0 <= 节点个数 <= 5000



注意：本题与主站 105 题重复：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package sword_offer

import (
	"testing"

	. "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_maxProfit(t *testing.T) {

}

func buildTree1(preorder []int, inorder []int) *TreeNode {

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

}
