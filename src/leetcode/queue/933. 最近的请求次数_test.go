/*
写一个 RecentCounter 类来计算最近的请求。

它只有一个方法：ping(int t)，其中 t 代表以毫秒为单位的某个时间。

返回从 3000 毫秒前到现在的 ping 数。

任何处于 [t - 3000, t] 时间范围之内的 ping 都将会被计算在内，包括当前（指 t 时刻）的 ping。

保证每次对 ping 的调用都使用比之前更大的 t 值。



示例：

输入：inputs = ["RecentCounter","ping","ping","ping","ping"], inputs = [[],[1],[100],[3001],[3002]]
输出：[null,1,2,3,3]


提示：

每个测试用例最多调用 10000 次 ping。
每个测试用例会使用严格递增的 t 值来调用 ping。
每次调用 ping 都有 1 <= t <= 10^9。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-recent-calls
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
	// . "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_ping(t *testing.T) {
	obj := Constructor()
	param_1 := obj.Ping(1)
	param_2 := obj.Ping(100)
	param_3 := obj.Ping(3001)
	param_4 := obj.Ping(3002)
	t.Log(param_1, param_2, param_3, param_4)
}

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}
func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	for this.queue[0] < t-3000 {
		this.queue = this.queue[1:]
	}
	return len(this.queue)
}
