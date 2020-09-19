/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
	// . "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_xxxxx(t *testing.T) {

}
func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}
	row := k / factorial(n-1)
	col := k % factorial(n-1)
	return string(row) + getPermutation(n-1, col)
}
func factorial(n int) int {
	var facVal = 1
	if n < 0 {

	} else {
		for i := 1; i <= n; i++ {
			facVal *= i
		}
	}
	return facVal

}
