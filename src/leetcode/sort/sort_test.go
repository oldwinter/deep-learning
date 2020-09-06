package sort

//
import (
	"sort"
	"testing"
)

// 350. 两个数组的交集 II
func intersect(nums1 []int, nums2 []int) []int {
	t := map[int]int{}
	result := []int{}
	for _, num := range nums1 {
		t[num]++
	}
	for _, num := range nums2 {
		if t[num] > 0 {
			result = append(result, num)
			t[num]--
		}
	}
	return result
}

// 349. 两个数组的交集
func intersection(nums1 []int, nums2 []int) []int {
	t := map[int]int{}
	result := []int{}
	for _, num := range nums1 {
		t[num] = 1
	}
	for _, num := range nums2 {
		if t[num] == 1 {
			result = append(result, num)
			t[num] = 0
		}
	}
	return result
}

// 905. 按奇偶排序数组
func sortArrayByParity(A []int) []int {
	result := make([]int, len(A))
	left := 0
	right := len(A) - 1
	for _, a := range A {
		if a%2 == 0 {
			result[left] = a
			left++
		} else {
			result[right] = a
			right--
		}
	}
	return result
}

// 922. 按奇偶排序数组 II
func sortArrayByParityII(A []int) []int {
	result := make([]int, len(A))
	left := 0
	right := len(A) - 1
	for _, a := range A {
		if a%2 == 0 {
			result[left] = a
			left += 2
		} else {
			result[right] = a
			right -= 2
		}
	}
	return result
}

// 976. 三角形的最大周长
func largestPerimeter(A []int) int {
	// sort.Ints(A)
	sort.Sort(sort.Reverse(sort.IntSlice(A)))
	for i := 0; i < len(A)-2; i++ {
		if A[i]-A[i+1] < A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}

type AAA struct {
	isExist bool
	index   int
}

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	a := map[int]int{}
	result := []int{}
	for i, num := range nums {
		if _, exist := a[target-num]; exist {
			result = append(result, a[target-num])
			result = append(result, a[num])

		}
		a[num] = i

	}
	return result
}

// 1560. 圆形赛道上经过次数最多的扇区
func mostVisited(n int, rounds []int) []int {
	return []int{}
}
func TestSort(t *testing.T) {
	if false {
		a := []int{1, 2, 2, 1}
		b := []int{2, 2}
		t.Log(intersect(a, b))

	}

	if false {
		a := []int{4, 2, 5, 7}
		// b := []int{2, 2}
		t.Log(sortArrayByParityII(a))

	}
	if false {
		a := []int{2, 1, 2}
		t.Log(largestPerimeter(a))
	}

	if true {
		a := []int{2, 7, 11, 15}
		t.Log(twoSum(a, 9))
	}
}
