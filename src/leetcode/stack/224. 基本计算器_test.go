/*
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格  。

示例 1:

输入: "1 + 1"
输出: 2
示例 2:

输入: " 2-1 + 2 "
输出: 3
示例 3:

输入: "(1+(4+5+2)-3)+(6+8)"
输出: 23
说明：

你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/basic-calculator
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
	"unicode"
)

func Test_calculate(t *testing.T) {
	t.Log(calculate("(1-3)-(6-8)"))
}
func calculate(s string) int {
	// numStack := []int{}
	operStack := []int{}
	// kuohaoStack := []byte{}
	var sum int = 0
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(rune(s[i])) {
			continue
		} else if s[i] == '(' || s[i] == ')' {
			continue
		} else if s[i] == '+' || s[i] == '-' {
			operStack = append(operStack, int(s[i]))
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
			var oper int = '+'
			if len(operStack) == 0 {
				oper = '+'
			} else {
				oper = operStack[len(operStack)-1]
			}

			if oper == '-' {
				sum -= num
			} else {
				sum += num
			}
			// }

		}

	}
	return int(sum)

}
