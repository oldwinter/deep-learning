package array

import (
	"testing"
)

func TestArrary(t *testing.T) {

	if true {
		in := []int{1, 2, 3}
		t.Log(countElements(in))
	}
	if false {
		nums := []int{5, 7, 1, 8}
		t.Log(checkPossibility(nums))
	}

	if false {
		nums := []int{1, 2, 3, 4}
		extraNum := 2
		t.Log(kidsWithCandies(nums, extraNum))
	}

	if false {
		nums := []int{1, 2, 3, 4}
		extraNum := 2
		t.Log(shuffle(nums, extraNum))
	}

	if false {
		nums := []int{1, 2, 3, 4}
		t.Log(runningSum(nums))

	}
}
