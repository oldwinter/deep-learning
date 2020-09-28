/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	t.Log(isValid("]]]"))
}
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
