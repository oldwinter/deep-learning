/*
给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deep-learning/leetcode/common"
)

func Test_zigzagLevelOrder(t *testing.T) {

}
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levelQ := []*TreeNode{root}
	result := [][]int{}
	levelCount := 0
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
		if levelCount%2 == 0 {
			result = append(result, tempI)
		} else {
			for i, j := 0, len(tempI)-1; i < j; {

				tempI[i], tempI[j] = tempI[j], tempI[i]
				i++
				j--
			}
			result = append(result, tempI)
		}
		levelCount++
		levelQ = tempQ

	}
	return result
}
