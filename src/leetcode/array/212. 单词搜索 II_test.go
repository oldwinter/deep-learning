/*
给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

示例:

输入:
words = ["oath","pea","eat","rain"] and board =
[
  ['o','a','a','n'],
  ['e','t','a','e'],
  ['i','h','k','r'],
  ['i','f','l','v']
]

输出: ["eat","oath"]
说明:
你可以假设所有输入都由小写字母 a-z 组成。

提示:

你需要优化回溯算法以通过更大数据量的测试。你能否早点停止回溯？
如果当前单词不存在于所有单词的前缀中，则可以立即停止回溯。什么样的数据结构可以有效地执行这样的操作？散列表是否可行？为什么？ 前缀树如何？如果你想学习如何实现一个基本的前缀树，请先查看这个问题： 实现Trie（前缀树）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-search-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
	// . "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_findWords(t *testing.T) {
	// in1 := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	// str1 := []string{"oath", "eat", "rain"}

	in2 := [][]byte{{'a', 'a'}}
	str2 := []string{"aaa"}
	t.Log(findWords(in2, str2))
}
func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 {
		return []string{}
	}
	if len(board[0]) == 0 {
		return []string{}
	}

	if len(words) == 0 {
		return []string{}
	}
	result := []string{}
	isVisitedMap := [][][]bool{}
	counts := []int{}
	for i := 0; i < len(words); i++ {
		temp := [][]bool{}
		for i := 0; i < len(board); i++ {
			temp = append(temp, make([]bool, len(board[0])))
		}
		isVisitedMap = append(isVisitedMap, temp)
	}

	for i := 0; i < len(words); i++ {
		counts = append(counts, 0)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			findWordsDFS(board, words, isVisitedMap, i, j, &counts, &result)
		}

	}
	return result
}

func findWordsDFS(board [][]byte, words []string, isVisited [][][]bool, x, y int, counts *[]int, result *[]string) {
	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
		return
	}
	for i := 0; i < len((words)); i++ {

		if isVisited[i][x][y] {
			return
		}

		if (*counts)[i] == len(words[i])-1 {
			// board[x][y] == words[0][(*counts)[0]]
			if board[x][y] == words[i][(*counts)[i]] {
				*result = append(*result, words[i])
				(*counts)[i] = -1
			}

			return
		}
		if (*counts)[i] != -1 && board[x][y] == words[i][(*counts)[i]] {
			(*counts)[i]++
			isVisited[i][x][y] = true
			findWordsDFS(board, words, isVisited, x+1, y, counts, result)
			findWordsDFS(board, words, isVisited, x-1, y, counts, result)
			findWordsDFS(board, words, isVisited, x, y+1, counts, result)
			findWordsDFS(board, words, isVisited, x, y-1, counts, result)
			isVisited[i][x][y] = false
		}

	}

	return

}
