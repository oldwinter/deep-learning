/*
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。



示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deep-learning/leetcode/common"
)

func Test_levelOrder(t *testing.T) {

}
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levelQ := []*TreeNode{root}
	result := [][]int{}
	for len(levelQ) != 0 {
		tempQ := []*TreeNode{}
		tempI := []int{}
		for _, v := range levelQ {
			tempI = append(tempI, v.Val)
			if v.Left != nil {
				tempQ = append(tempQ, v.Left)
			}

			if v.Right != nil {
				tempQ = append(tempQ, v.Right)
			}

		}
		result = append(result, tempI)
		levelQ = tempQ

	}
	return result
}
