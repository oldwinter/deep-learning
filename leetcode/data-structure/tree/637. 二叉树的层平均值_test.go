/*
给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。



示例 1：

输入：
    3
   / \
  9  20
    /  \
   15   7
输出：[3, 14.5, 11]
解释：
第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。


提示：

节点值的范围在32位有符号整数范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deep-learning/leetcode/common"
)

func Test_averageOfLevels(t *testing.T) {

}
func averageOfLevels(root *TreeNode) []float64 {

	if root == nil {
		return []float64{}
	}
	result := []float64{}
	levelQ := []*TreeNode{root}
	for len(levelQ) != 0 {
		nextQ := []*TreeNode{}
		tempVal := 0.0
		for _, v := range levelQ {

			tempVal += float64(v.Val)
			if v.Left != nil {
				nextQ = append(nextQ, v.Left)
			}
			if v.Right != nil {
				nextQ = append(nextQ, v.Right)
			}

		}
		result = append(result, tempVal/float64(len(levelQ)))
		levelQ = nextQ

	}

	return result
}
