/*
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_levelOrderBottom(t *testing.T) {

}

func levelOrderBottom(root *TreeNode) [][]int {
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
	for i := 0; i < len(results)/2; i++ {
		results[i], results[len(results)-i-1] = results[len(results)-i-1], results[i]
	}

	return results
}
