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
