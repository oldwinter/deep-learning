/*
给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。



示例:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
	// . "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_exist(t *testing.T) {
	t.Log(exist([][]byte{{'a', 'a', 'a'}}, "aaaa"))
}

// 79. 单词搜索
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	if len(board[0]) == 0 {
		return false
	}

	if word == "" {
		return true
	}

	isVisitedMap := [][]bool{}
	for i := 0; i < len(board); i++ {
		isVisitedMap = append(isVisitedMap, make([]bool, len(board[0])))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if existDFS(board, word, isVisitedMap, i, j, 0) {
				return true
			}
		}

	}
	return false

}

func existDFS(board [][]byte, word string, isVisited [][]bool, x, y, count int) bool {

	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
		return false
	}
	if isVisited[x][y] {
		return false
	}

	if count == len(word)-1 {
		return word[count] == board[x][y]
	}
	if word[count] == board[x][y] {
		isVisited[x][y] = true
		up := existDFS(board, word, isVisited, x+1, y, count+1)
		if up {
			return true
		}
		// isVisited[x][y] = false
		down := existDFS(board, word, isVisited, x-1, y, count+1)
		if down {
			return true
		}
		// isVisited[x][y] = false
		left := existDFS(board, word, isVisited, x, y-1, count+1)
		if left {
			return true
		}
		// isVisited[x][y] = false
		right := existDFS(board, word, isVisited, x, y+1, count+1)
		if right {
			return true
		}
		isVisited[x][y] = false
	}

	return false
	// // isVisited[x][y]=false
	// return up || down || left || right
}
