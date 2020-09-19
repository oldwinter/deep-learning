/*
给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。



示例 1：

输入：[1,2,3]

       1
      / \
     2   3

输出：6
示例 2：

输入：[-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出：42

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-maximum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_maxPathSum(t *testing.T) {

}
func maxPathSum(root *TreeNode) int {
	max1, max2 := maxPathSumDFS(root)
	if max1 > max2 {
		return max1
	} else {
		return max2
	}
}

//max 代表左右都包含的子树
func maxPathSumDFS(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	maxInner := root.Val
	maxOuter := root.Val

	maxInnerLeft, MaxOuterLeft := maxPathSumDFS(root.Left)
	maxInnerRight, MaxOuterRight := maxPathSumDFS(root.Right)
	if MaxOuterLeft > 0 && MaxOuterRight > 0 {
		maxInner = MaxOuterLeft + MaxOuterRight + root.Val
		if maxInner < maxInnerLeft {
			maxInner = maxInnerLeft
		}
		if maxInner < maxInnerRight {
			maxInner = maxInnerRight
		}
		if MaxOuterRight > MaxOuterLeft {
			maxOuter = root.Val + MaxOuterRight
		} else {
			maxOuter = root.Val + MaxOuterLeft
		}
	} else if MaxOuterLeft > 0 && MaxOuterRight <= 0 {
		maxOuter = root.Val + MaxOuterLeft
		if root.Val > 0 {
			maxInner = maxOuter
		} else {
			maxInner = maxInnerLeft
		}
	} else if MaxOuterLeft <= 0 && MaxOuterRight > 0 {
		maxOuter = root.Val + MaxOuterRight
		if root.Val > 0 {
			maxInner = maxOuter
		} else {
			maxInner = maxInnerRight
		}
	} else {
		if maxInnerLeft < maxInnerRight {
			maxInner = maxInnerLeft
		} else {
			maxInner = maxInnerRight
		}
		maxOuter = maxInner
	}

	return maxInner, maxOuter
}
