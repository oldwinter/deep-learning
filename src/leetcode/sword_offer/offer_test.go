package sword_offer

import (
	. "leetcode/src/leetcode/linkedlist"
	. "leetcode/src/leetcode/tree"
	"testing"
)

func TestOffer(t *testing.T) {

	if false {
		t.Log(add(5, 7))
	}

	if false {
		t.Log(isStraight([]int{0, 0, 8, 5, 4}))
	}

	if false {
		t.Log((twoSum111(3)))
	}
	if false {
		t.Log(reverseWords(" hello world!"))
	}

	if false {
		t.Log(findContinuousSequence(15))
	}

	if false {
		in := []int{5, 7, 7, 8, 8, 10}

		t.Log(search(in, 8))
	}

	if false {
		in := []int{5, 4, 11, 7, 2, 8, 13, 3, 6, 1}
		t.Log(getLeastNumbers(in, 3))
	}

	if false {

		in := &Node{Val: 1, Next: nil, Random: nil}
		in2 := &Node{Val: 2, Next: nil, Random: nil}
		in.Next = in2
		in2.Random = in2

		t.Log(copyRandomList(in))
	}
	if false {
		preorder := []int{5, 4, 11, 7, 2, 8, 13, 3, 6, 1}
		inorder := []int{7, 11, 2, 4, 5, 13, 8, 6, 3, 1}

		in := buildTree(preorder, inorder)
		t.Log(pathSum(in, 22))
	}

	if false {
		in := []int{5, 4, 3, 2, 1}
		t.Log(verifyPostorder(in))
	}
	if false {
		in := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
		t.Log(spiralOrder(in))
	}
	if false {
		t.Log(hammingWeight(3))
	}
	if false {
		t.Log(movingCount(11, 8, 16))
	}

	if false {
		// in := [][]byte{
		// 	{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
		// }
		in := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
		t.Log(exist(in, "ABCESEEEFS"))
	}
	if false {
		in := []int{1, 3}
		t.Log(minArray(in))
	}

	if false {
		t.Log(fib(8))
	}

	if false {
		obj := Constructor()
		obj.DeleteHead()
		obj.AppendTail(5)
		obj.AppendTail(2)
		obj.DeleteHead()
		obj.DeleteHead()
	}

	if false {
		preorder := []int{3, 9, 20, 15, 7}
		inorder := []int{9, 3, 15, 20, 7}
		result := buildTree(preorder, inorder)
		t.Log(PrintTree(result))
	}

	if false {
		inList := MakeLinkedList([]int{1, 3, 2})
		t.Log(reversePrint(inList))
	}

	if false {
		t.Log(replaceSpace("We are happy."))
	}

	if false {
		// in := [][]int{
		// 	{1, 4, 7, 11, 15},
		// 	{2, 5, 8, 12, 19},
		// 	{3, 6, 9, 16, 22},
		// 	{10, 13, 14, 17, 24},
		// 	{18, 21, 23, 26, 30},
		// }
		in2 := [][]int{
			{5},
		}
		t.Log(findNumberIn2DArray(in2, 5))
	}

	if false {
		in := []int{2, 3, 1, 0, 2, 5, 3}
		t.Log(findRepeatNumber(in))
	}

}
