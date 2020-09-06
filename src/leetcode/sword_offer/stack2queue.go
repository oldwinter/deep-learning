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

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
