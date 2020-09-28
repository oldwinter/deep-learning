/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )



示例 1：

输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]
示例 2：

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package sword_offer

type CQueue struct {
	stackA []int //进数
	stackB []int //从A进数，出数
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stackA = append(this.stackA, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackB) == 0 {
		for i := len(this.stackA) - 1; i > -1; i-- {
			this.stackB = append(this.stackB, this.stackA[i])
		}
		this.stackA = []int{}
		result := 0
		if len(this.stackB) == 0 {
			result = -1
			return result
		} else {
			result = this.stackB[len(this.stackB)-1]
		}

		this.stackB = this.stackB[:len(this.stackB)-1]
		return result
	} else {
		result := this.stackB[len(this.stackB)-1]
		this.stackB = this.stackB[:len(this.stackB)-1]
		return result
	}

}
