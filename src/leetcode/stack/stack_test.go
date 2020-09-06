package stack

// leetcode上关于栈的题目大家可以先做20,155,232,844,224,682,496
import (
	"fmt"
	"strconv"
	"testing"
	"unicode"
)

//20. 有效的括号
func isValid(s string) bool {

	var stack []byte
	var pairMap = map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {

			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				stack = append(stack, s[i])
			}
			if stack[len(stack)-1] == pairMap[s[i]] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, s[i])
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}

}

//155. 最小栈
type MinStack struct {
	stack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return 0
	}
	min := this.stack[0]
	for _, v := range this.stack {
		if v < min {
			min = v
		}
	}
	return min
}

//232. 用栈实现队列
type MyQueue struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructorssss() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	result := this.queue[0]
	this.queue = this.queue[1:]
	return result
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.queue[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.queue) == 0 {
		return true
	}
	return false
}

// 844. 比较含退格的字符串
func backspaceCompare(S string, T string) bool {

	Safter := string(getStack(S))
	Tafter := string(getStack(T))
	if Safter == Tafter {
		return true
	}
	return false

}
func getStack(S string) []byte {
	Sstack := []byte{}
	for i := 0; i < len(S); i++ {
		if S[i] != '#' {
			Sstack = append(Sstack, S[i])
		} else {
			if len(Sstack) == 0 {
				continue
			} else {
				Sstack = Sstack[:len(Sstack)-1]
			}
		}
	}
	return Sstack
}

//224. 基本计算器
func calculate(s string) int {
	numStack := []int{}
	operStack := []rune{}
	// kuohaoStack := []byte{}
	var sum int = 0
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i])) {
			continue
		} else if s[i] == ')' {
			// operStack = append(operStack, int(s[i]))

		} else if s[i] == '+' || s[i] == '-' || s[i] == '(' || s[i] == ')' {
			operStack = append(operStack, rune(s[i]))
		} else if s[i] >= '0' && s[i] <= '9' {
			// if len(operStack) != 0 {

			num := 0
			for ; i < len(s); i++ {
				if s[i] < '0' || s[i] > '9' {
					break
				}
				num = num*10 + int(s[i]-'0')
			}
			i--

			numStack = append(numStack, num)

			var oper rune = '+'
			if len(operStack) == 0 {
				oper = '+'
			} else {
				oper = operStack[len(operStack)-1]
			}

			if oper == '-' {
				sum -= numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
			} else if oper == '+' {
				sum += numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
			} else if oper == '(' {
				sum += numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
			} else if oper == ')' {
				oper = operStack[len(operStack)-1]
				operStack = operStack[:len(numStack)-1]
			}
			// }

		}

	}
	return int(sum)

}

// 682. 棒球比赛
func calPoints(ops []string) int {
	sum := 0
	sumStack := []int{}
	for _, op := range ops {
		if op == "+" {
			num := sumStack[len(sumStack)-1] + sumStack[len(sumStack)-2]
			sumStack = append(sumStack, num)
		} else if op == "D" {
			num := sumStack[len(sumStack)-1] * 2
			sumStack = append(sumStack, num)
		} else if op == "C" {
			sumStack = sumStack[:len(sumStack)-1]
		} else {
			num, err := strconv.Atoi(op)
			if err != nil {
				fmt.Print("failed")
			}
			sumStack = append(sumStack, num)
		}
	}
	for _, v := range sumStack {
		sum += v
	}
	return sum
}

// 496. 下一个更大元素 I
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
func TestLinkedlist(t *testing.T) {
	if false {
		t.Log(isValid("]]]"))
	}

	if false {
		obj := Constructor()
		obj.Push(1)
		obj.Push(3)
		obj.Push(4)
		obj.Pop()
		param3 := obj.Top()
		param4 := obj.GetMin()
		t.Log(param3, param4)
	}

	if false {

		obj := Constructorssss()
		obj.Push(2)
		obj.Push(3)
		param_2 := obj.Pop()
		param_3 := obj.Peek()
		param_4 := obj.Empty()
		t.Log(param_2, param_3, param_4)

	}

	if false {
		t.Log(backspaceCompare("a##c", "#a#cc"))
	}

	//未成功，太难
	if false {
		t.Log(calculate("(1-3)-(6-8)"))
	}

	if false {
		s := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
		t.Log(calPoints(s))
	}

	if true {
		nums1 := []int{1, 3, 5, 2, 4}
		nums2 := []int{6, 5, 4, 3, 2, 1, 7}
		t.Log(nextGreaterElement(nums1, nums2))
	}

}
