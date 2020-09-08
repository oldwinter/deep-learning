package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先遍历，以数组形式存放
// 关键点：用一个队列，把树的层的前后给记住
func PrintTree(head *TreeNode) []int {
	if head == nil {
		return []int{}
	}
	result := []int{}
	a := []TreeNode{}
	a = append(a, *head)
	for len(a) != 0 {
		head := a[0]
		result = append(result, head.Val)
		a = a[1:]
		if head.Left != nil {
			a = append(a, *head.Left)
			// result = append(result, head.Left.Val)
		} else {
			// result = append(result, -1)
		}
		if head.Right != nil {
			a = append(a, *head.Right)
		} else {
			// result = append(result, -1)
		}

	}
	return result
}

func PrintTree2(root *TreeNode) [][]int {
	results := [][]int{}
	if root == nil {
		return results
	}
	levelQueue := []TreeNode{*root}
	level := 0
	for len(levelQueue) != 0 {
		a := []TreeNode{}
		results = append(results, []int{})
		for _, v := range levelQueue {
			// 取值
			results[level] = append(results[level], v.Val)
			// 入队
			if v.Left != nil {
				a = append(a, *v.Left)
			}
			if v.Right != nil {
				a = append(a, *v.Right)
			}

		}
		level++
		levelQueue = a

	}

	return results
}

// 107. 二叉树的层次遍历 II
func levelOrderBottom(root *TreeNode) [][]int {
	results := [][]int{}
	if root == nil {
		return results
	}
	levelQueue := []TreeNode{*root}
	level := 0
	for len(levelQueue) != 0 {
		a := []TreeNode{}

		results = append(results, []int{})
		for _, v := range levelQueue {
			// 取值
			results[level] = append(results[level], v.Val)
			// 入队
			if v.Left != nil {
				a = append(a, *v.Left)
			}
			if v.Right != nil {
				a = append(a, *v.Right)
			}

		}
		level++
		levelQueue = a

	}
	for i := 0; i < len(results)/2; i++ {
		results[i], results[len(results)-i-1] = results[len(results)-i-1], results[i]
	}

	return results
}

// 108. 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	root := TreeNode{Val: 0, Left: nil, Right: nil}
	if len(nums) == 0 {
		return nil
	}
	// if len(nums)==1 {
	// 	root.Val = nums[0]
	// 	return &root

	// }
	root.Val = nums[len(nums)/2]
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	return &root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  538. 把二叉搜索树转换为累加树

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	// if root.Right != nil {
	convertBSTRecur(root, &sum)

	return root
}

func convertBSTRecur(root *TreeNode, sum *int) int {
	convertBSTRecur(root.Right, sum)
	*sum += root.Val
	root.Val += *sum
	convertBSTRecur(root.Left, sum)
	return root.Val
}

// 543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// result := 0
	leftLength := 0
	rightLength := 0
	if root.Left != nil {
		leftLength = diameterOfBinaryTree(root.Left) + 1
	}
	if root.Right != nil {
		rightLength = diameterOfBinaryTree(root.Right) + 1
	}

	return leftLength + rightLength

}
