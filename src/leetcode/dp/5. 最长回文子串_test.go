/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("babad"))
}
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start := 0
	maxlen := 1
	ss := make([][]bool, len((s)))
	for i := 0; i < len(s); i++ {
		ss[i] = make([]bool, len(s))
		ss[i][i] = true
	}
	for i := 0; i < len(s); i++ {

		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && i+1 >= 0 && j-1 >= 0 && i+1 < len(s) && j-1 < len(s) && ss[i+1][j-1] {
				ss[i][j] = true
			} else {
				ss[i][j] = false
			}
			if ss[i][j] {
				if j-i+1 > maxlen {
					maxlen = j - i + 1
					start = i
				}
			}

		}
	}
	return s[start : start+maxlen]

}
