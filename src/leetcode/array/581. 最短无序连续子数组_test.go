/*
给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

你找到的子数组应是最短的，请输出它的长度。

示例 1:

输入: [2, 6, 4, 8, 10, 9, 15]
输出: 5
解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
说明 :

输入的数组长度范围在 [1, 10,000]。
输入的数组可能包含重复元素 ，所以升序的意思是<=。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"sort"
	"testing"
	// . "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_findUnsortedSubarray(t *testing.T) {

}
func findUnsortedSubarray(nums []int) int {
	temp := []int{}
	temp = append(temp, nums...)

	sort.Ints(temp)
	left := 0
	right := 0
	for i := 0; i < len(nums); i++ {
		if temp[i] != nums[i] {
			left = i
			break
		}
	}
	for j := len(nums) - 1; j > 0; j-- {
		if temp[j] != nums[j] {
			right = j
			break
		}
	}
	if left == 0 && right == 0 {
		return 0
	}
	return right - left + 1
}
