/*

 */
package leetcode

import (
	"testing"

	. "github.com/oldwinter/deep-learning/leetcode/common"
)

func Test_sortedArrayToBST(t *testing.T) {

}
func sortedArrayToBST(nums []int) *TreeNode {
	root := TreeNode{Val: 0, Left: nil, Right: nil}
	if len(nums) == 0 {
		return nil
	}
	root.Val = nums[len(nums)/2]
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	return &root
}
