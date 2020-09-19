/*
给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。



示例 1:

输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
    对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
    对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
示例 2:

输入: nums1 = [2,4], nums2 = [1,2,3,4].
输出: [3,-1]
解释:
    对于 num1 中的数字 2 ，第二个数组中的下一个较大数字是 3 。
    对于 num1 中的数字 4 ，第二个数组中没有下一个更大的数字，因此输出 -1 。


提示：

nums1和nums2中所有元素是唯一的。
nums1和nums2 的数组大小都不超过1000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-greater-element-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	nums1 := []int{1, 3, 5, 2, 4}
	nums2 := []int{6, 5, 4, 3, 2, 1, 7}
	t.Log(nextGreaterElement(nums1, nums2))
}
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	if len(nums1) == 0 {
		return []int{}
	}
	if len(nums2) == 0 {
		return []int{}
	}
	newNum := map[int]int{}

	for i := 0; i < len(nums2)-1; i++ {
		k := i
		for k < len(nums2)-1 && nums2[k+1] < nums2[i] {
			k++
		}
		if k < len(nums2)-1 {
			newNum[nums2[i]] = nums2[k+1]
		} else {
			newNum[nums2[i]] = -1
		}
	}
	newNum[nums2[len(nums2)-1]] = -1

	result := []int{}
	for _, num := range nums1 {
		result = append(result, newNum[num])
	}
	return result
}
