/*
给定两个数组，编写一个函数来计算它们的交集。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
	// . "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_intersection(t *testing.T) {

}

// 349. 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	t := map[int]int{}
	result := []int{}
	for _, num := range nums1 {
		t[num] = 1
	}
	for _, num := range nums2 {
		if t[num] == 1 {
			result = append(result, num)
			t[num] = 0
		}
	}
	return result
}
