package stack

//
import (
	"testing"
)

// 933. 最近的请求次数
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

// 剑指 Offer 59 - I. 滑动窗口的最大值 ；同 239. 滑动窗口最大值
// func maxSlidingWindow(nums []int, k int) []int {
// 	if len(nums) == 0 {
// 		return []int{}
// 	}
// 	result := []int{}
// 	queue := []int{}
// 	for i := 0; i < k; i++ {
// 		queue = append(queue, nums[i])
// 	}
// 	for i := 0; i < len(nums)-k+1; i++ {
// 		max := nums[i]
// 		for _, q := range queue {
// 			if q > max {
// 				max = q
// 			}
// 		}
// 		result = append(result, max)
// 		if i+k < len(nums) {
// 			queue = append(queue, nums[i+k])
// 			queue = queue[1:]
// 		}

// 	}
// 	return result
// }
// 降低queqe中找最大值的时间复杂度 从维护一个滑动窗口队列，变成维护一个有序递减队列
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := []int{}
	queue := []int{0}
	queue[0] = nums[0]
	for i := 1; i < k; i++ {
		if nums[i] > queue[len(queue)-1] {
			queue = append(queue, nums[i])
		} else {
			for nums[i] > queue[0] {
				queue = queue[1:]
			}
			if nums[i] < queue[0] {
				temp := []int{nums[i]}
				queue = append(temp, queue...)
			}
		}

	}
	// max := nums[i]
	// 	for _, q := range queue {
	// 		if q > max {
	// 			max = q
	// 		}
	// 	}
	// queue[0] = nums[0]
	result = append(result, queue[len(queue)-1])
	for i := k; i < len(nums); i++ {
		// queue=append(queue,nums[i])
		if nums[i] > queue[len(queue)-1] {
			queue = append(queue, nums[i])
		} else {
			for nums[i] > queue[0] {
				queue = queue[1:]
			}
			for nums[i] < queue[0] {
				temp := []int{nums[i]}
				queue = append(temp, queue...)
				if len(queue) > k {
					queue = queue[:len(queue)-1]
				}
			}
		}

		result = append(result, queue[len(queue)-1])

	}
	return result
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

// 作者：Rolle
// 链接：https://leetcode-cn.com/problems/get-kth-magic-number-lcci/solution/go-dp-by-rolle/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func TestQueue(t *testing.T) {
	if false {
		obj := Constructor()
		param_1 := obj.Ping(1)
		param_2 := obj.Ping(100)
		param_3 := obj.Ping(3001)
		param_4 := obj.Ping(3002)
		t.Log(param_1, param_2, param_3, param_4)

	}

	if false {
		nums := []int{1, -1}
		t.Log(maxSlidingWindow(nums, 1))
	}
	if true {
		t.Log(getKthMagicNumber(9))
	}
}
