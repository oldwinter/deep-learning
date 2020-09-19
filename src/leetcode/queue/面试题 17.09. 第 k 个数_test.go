/*
有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。

示例 1:

输入: k = 5

输出: 9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/get-kth-magic-number-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
	// . "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_getKthMagicNumber(t *testing.T) {
	t.Log(getKthMagicNumber(9))
}
func getKthMagicNumber(k int) int {
	queue := make([]int, k)
	queue[0] = 1
	a := 0
	b := 0
	c := 0
	for j := 1; j < k; j++ {
		queue[j] = min(queue[a]*3, queue[b]*5, queue[c]*7)
		// queue = append(queue, key)
		if queue[j] == queue[a]*3 {
			a++
		}
		if queue[j] == queue[b]*5 {
			b++
		}
		if queue[j] == queue[c]*7 {
			c++
		}
		// queue = append(queue, key)

	}

	return queue[k-1]

}
func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}

	if b <= a && b <= c {
		return b
	}

	return c
}
