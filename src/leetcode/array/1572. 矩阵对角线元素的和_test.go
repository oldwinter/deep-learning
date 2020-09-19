/*
给你一个正方形矩阵 mat，请你返回矩阵对角线元素的和。

请你返回在矩阵主对角线上的元素和副对角线上且不在主对角线上元素的和。



示例  1：



输入：mat = [[1,2,3],
            [4,5,6],
            [7,8,9]]
输出：25
解释：对角线的和为：1 + 5 + 9 + 3 + 7 = 25
请注意，元素 mat[1][1] = 5 只会被计算一次。
示例  2：

输入：mat = [[1,1,1,1],
            [1,1,1,1],
            [1,1,1,1],
            [1,1,1,1]]
输出：8
示例 3：

输入：mat = [[5]]
输出：5


提示：

n == mat.length == mat[i].length
1 <= n <= 100
1 <= mat[i][j] <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/matrix-diagonal-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"testing"
	// . "github.com/oldwinter/deepLearning/src/leetcode/common"
)

func Test_diagonalSum(t *testing.T) {

}

func diagonalSum(mat [][]int) int {
	// result:=0
	result1 := 0
	result2 := 0
	result3 := 0
	for i := 0; i < len(mat); i++ {
		result1 += mat[i][i]
	}
	for j := 0; j < len(mat); j++ {
		result2 += mat[j][len(mat)-j-1]
	}
	if len(mat)%2 == 0 {
		result3 = 0
	} else {
		result3 = mat[len(mat)/2][len(mat)/2]
	}
	return result1 + result2 - result3
}
