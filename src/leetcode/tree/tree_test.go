package tree

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func sortedArrayToBST(nums []int) *TreeNode {

// }

// 17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	charMaps := map[string]string{
		"1": "",
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	fmt.Print(charMaps)
	result := []string{}

	return result
}

func TestLinkedlist(t *testing.T) {

	if true {
		t.Log(letterCombinations("23"))
	}
	if false {
		// in := []int{-10, -3, 0, 5, 9}
		// t.Log(sortedArrayToBST(in))

	}
}
