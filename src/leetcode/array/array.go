package array

// 665. 非递减数列
func checkPossibility1(nums []int) bool {
	flag := 0
	prevPrevNum := 0
	prevNum := 0

	for i, num := range nums {
		if i == 0 {
			prevNum = num
		} else if i == 1 {
			prevPrevNum = prevNum
			prevNum = num
		} else {
			if num < prevNum && num < prevPrevNum {
				flag++
				prevNum = prevPrevNum
			} else if num < prevNum && num >= prevPrevNum {
				prevPrevNum = prevNum
				prevNum = num
			} else {
				prevPrevNum = prevNum
				prevNum = num
			}

		}

	}

	return flag < 1
}

func checkPossibility(nums []int) bool {
	flag := 0
	for i := 0; i < len(nums); i++ {
		if i > 2 {
			if nums[i] < nums[i-1] && nums[i-1] < nums[i-2] {
				flag++
			}
		}
	}

	return flag < 1
}

// 1431. 拥有最多糖果的孩子
func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := []bool{}
	maxNum := 0
	for _, candi := range candies {
		if candi > maxNum {
			maxNum = candi
		}
	}
	for _, candi := range candies {
		if candi < maxNum-extraCandies {
			result = append(result, false)
		} else {
			result = append(result, true)
		}
	}
	return result
}

// 1470. 重新排列数组
func shuffle(nums []int, n int) []int {
	result := []int{}
	for i := 0; i < n; i++ {
		result = append(result, nums[i])
		result = append(result, nums[n+i])
	}
	return result
}

// 1480. 一维数组的动态和
func runningSum(nums []int) []int {
	numNew := []int{}
	for i, num := range nums {
		if i > 0 {
			numNew = append(numNew, numNew[i-1]+num)
		} else {
			numNew = append(numNew, num)
		}

	}
	return numNew
}

// 5491. 矩阵对角线元素的和
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

// 1426. 数元素
func countElements(arr []int) int {

	m := map[int]int{}
	result := 0
	for _, a := range arr {
		m[a]++
	}
	for k := range m {
		if m[k+1] > 0 {
			result += m[k]
		}

	}
	return result
}

// 53. 最大子序和
func maxSubArray(nums []int) int {
	// result :=0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}
		if i > 0 && nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		} else {
			continue
		}
	}
	result := nums[0]
	for _, num := range nums {
		if num > result {
			result = num
		}
	}
	return result
}

// 69. x 的平方根
func mySqrt(x int) int {
	left := 0
	right := x
	for left < right {
		mid := (left + right) / 2
		if mid*mid > x {
			right = mid - 1
		} else if mid*mid < x {
			left = mid + 1
		} else {
			return mid
		}

	}
	if left*left > x {
		return left - 1
	} else {
		return left
	}

}

// 268. 缺失数字
func missingNumber(nums []int) int {

}
