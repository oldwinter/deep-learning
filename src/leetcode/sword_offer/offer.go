package sword_offer

import (
	"container/heap"
	"math"
	"regexp"
	"sort"
	"strings"
	"unicode"

	. "github.com/oldwinter/deepLearning/src/leetcode/linkedlist"
	. "github.com/oldwinter/deepLearning/src/leetcode/tree"
)

// 剑指 Offer 03. 数组中重复的数字
func findRepeatNumber(nums []int) int {
	m := map[int]int{}

	for _, num := range nums {
		if m[num] > 0 {
			return num
		} else {
			m[num]++
		}
	}
	return -1
}

// 剑指 Offer 04. 二维数组中的查找
// 关键点：要从右上角或者左下角开始遍历 类似于取堆的堆顶，小于堆顶取左子树，大于堆顶取右子树
func findNumberIn2DArray(matrix [][]int, target int) bool {

	for i, j := len(matrix)-1, 0; i > -1 && j < len(matrix[0]); {
		if target > matrix[i][j] {
			j++
		} else if target < matrix[i][j] {
			i--
		} else {
			return true
		}
	}
	return false

}

// 剑指 Offer 05. 替换空格
// 要点，%20是3个字符，不是一个字符，需要append 3次
func replaceSpace(s string) string {
	result := []rune{}
	for _, c := range s {
		if unicode.IsSpace(c) {
			result = append(result, rune('%'))
			result = append(result, rune('2'))
			result = append(result, rune('0'))
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}

// 剑指 Offer 06. 从尾到头打印链表
func reversePrint(head *ListNode) []int {

	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}

// 剑指 Offer 07. 重建二叉树 假设不含重复数字
// 105. 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	// result := TreeNode{}
	// for  {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}

	if len(preorder) == 1 {
		return root
	}

	mid := 0
	for i, v := range inorder {
		if v == preorder[0] {
			mid = i
		}
	}
	mid++

	leftnode := buildTree(preorder[1:mid], inorder[:mid-1])
	rightnode := buildTree(preorder[mid:], inorder[mid:])

	root.Left = leftnode
	root.Right = rightnode
	return root
	// }
}

// 剑指 Offer 10- I. 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := []int{0, 1}
	for i := 2; i < n+1; i++ {
		temp := a[i-1] + a[i-2]
		a = append(a, temp%(1e9+7))
	}
	return int(a[n])
}

// 剑指 Offer 10- II. 青蛙跳台阶问题
func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	a := []int{1, 1}
	for i := 2; i < n+1; i++ {
		temp := a[i-1] + a[i-2]
		a = append(a, temp%(1e9+7))
	}
	return int(a[n])
}

// 剑指 Offer 11. 旋转数组的最小数字
// 关键点：发现也可以二分法 排序数组的查找，第一个应该想到的是二分法。
// 二分法关键点，瞄准左右区间，每个循环去缩小区间，直到区间为1
func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	//顺序遍历，但可以找到即停
	// for i := 1; i < len(numbers); i++ {
	// 	if numbers[i] < numbers[i-1] {
	// 		return numbers[i]
	// 	}
	// }
	// return numbers[0]

	//二分查找
	i, k := 0, len(numbers)-1
	for i < k {
		j := (i + k) / 2
		if numbers[i] < numbers[k] {
			return numbers[i]
		}
		if numbers[j] < numbers[k] {
			k = j
			// j = (i + j) / 2

			// i = 0+i/2
		} else if numbers[j] > numbers[k] {
			i = j + 1
			// j = (j + k) / 2

		} else {
			k--
		}
	}

	return numbers[i]

}

// 剑指 Offer 12. 矩阵中的路径
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {

			if dfsSearch(board, word, i, j) {
				return true
			}

		}
	}

	return false
}

// in: 当前图，当前点，待匹配串
// out: 是否找到匹配串
func dfsSearch(board [][]byte, word string, x, y int) bool {
	if len(word) == 0 {
		return true
	}
	if x < 0 || y < 0 || x > len(board)-1 || y > len(board[0])-1 {
		return false
	}

	if board[x][y] == word[0] {
		temp := board[x][y]
		board[x][y] = '-'
		subResult := dfsSearch(board, word[1:], x-1, y) || dfsSearch(board, word[1:], x+1, y) || dfsSearch(board, word[1:], x, y-1) || dfsSearch(board, word[1:], x, y+1)
		if !subResult {
			board[x][y] = temp
		}
		return subResult
	} else {

		return false
	}

}

// 剑指 Offer 13. 机器人的运动范围
// DFS 递归  BFS 队列
func movingCount(m int, n int, k int) int {
	isVisitedMap := map[int]bool{}
	return dfsSearchBlock(0, 0, k, m, n, isVisitedMap)

}

func dfsSearchBlock(i, j, k, m, n int, maps map[int]bool) int {

	if i < 0 || j < 0 || i > m-1 || j > n-1 {
		return 0
	}

	if bitSum(i)+bitSum(j) <= k && !maps[i*100+j] {
		maps[i*100+j] = true
		num := 1 + dfsSearchBlock(i+1, j, k, m, n, maps) + dfsSearchBlock(i, j+1, k, m, n, maps)

		return num
	} else {
		return 0
	}

}

func bitSum(in int) int {
	sum := 0
	for in != 0 {
		sum += in % 10
		in = in / 10
	}
	return sum
}

// 剑指 Offer 13. 机器人的运动范围
// DFS 递归  BFS 队列
func movingCountBFS(m int, n int, k int) int {
	isVisitedMap := map[int]bool{}
	return dfsSearchBlock(0, 0, k, m, n, isVisitedMap)

}

// 剑指 Offer 15. 二进制中1的个数
func hammingWeight(num uint32) int {
	sum := 0
	for num != 0 {
		sum += int(num % 2)
		num = num / 2
	}
	return sum
}

// 剑指 Offer 17. 打印从1到最大的n位数
func printNumbers(n int) []int {

	result := []int{}
	for i := 0; i < int(math.Pow(10, float64(n)))-1; i++ {
		result = append(result, i+1)
	}
	return result
}

// 剑指 Offer 18. 删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	result := &ListNode{}
	result.Next = head
	prev := result
	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
			break
		}
		prev = head
		head = head.Next
	}
	return result.Next
}

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 != 1 {
			for nums[right]%2 != 1 {
				right--
				if left >= right {
					break
				}
			}
			nums[left], nums[right] = nums[right], nums[left]
		}
		left++
	}
	return nums
}

// 剑指 Offer 22. 链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	left := &ListNode{}
	left = head
	right := &ListNode{}
	right = head
	for i := 0; i < k; i++ {
		right = right.Next
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}
	return left
}

// 剑指 Offer 24. 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev := &ListNode{}
	prev = nil
	for head != nil {
		temp := head.Next

		head.Next = prev
		prev = head
		head = temp

	}

	return prev
}

// 剑指 Offer 25. 合并两个排序的链表
// 关键优化点，剩余的l1或l2 已经是有序，直接附加在pointer后面即可
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	guard := &ListNode{}
	pointer := guard
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pointer.Next = l1
			l1 = l1.Next
		} else {
			pointer.Next = l2
			l2 = l2.Next
		}
		pointer = pointer.Next
	}
	if l1 != nil {
		pointer.Next = l1
	} else {
		pointer.Next = l2
	}

	// for l1 != nil {
	// 	pointer.Next = l1
	// 	l1 = l1.Next
	// 	pointer = pointer.Next
	// }

	// for l2 != nil {
	// 	pointer.Next = l2
	// 	l2 = l2.Next
	// 	pointer = pointer.Next
	// }
	return guard.Next
}

// 剑指 Offer 27. 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = mirrorTree(root.Left), mirrorTree(root.Right)
	return root
}

// 剑指 Offer 28. 对称的二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return digui(root.Left, root.Right)
}
func digui(left, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return digui(left.Left, right.Right) && digui(left.Right, right.Left)

}

// 剑指 Offer 29. 顺时针打印矩阵
// 一定是按照右下左上的顺序 ，遇到墙壁就调整方向，边界变化
// 最后一轮有多余遍历，把数据裁剪掉即可
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	// visited := map[int]bool{}
	result := []int{}
	// 0,1,2,3  代表右下左上的顺序
	rightRange := len(matrix[0]) - 1
	leftRange := 0
	topRange := 0
	bottomRange := len(matrix) - 1
	i, j := 0, 0
	for topRange <= bottomRange && leftRange <= rightRange {

		//左-右

		for j = leftRange; j <= rightRange; j++ {
			result = append(result, matrix[topRange][j])
		}
		topRange++
		//左-右
		for i = topRange; i <= bottomRange; i++ {
			result = append(result, matrix[i][rightRange])
		}
		rightRange--

		for j = rightRange; j >= leftRange; j-- {
			result = append(result, matrix[bottomRange][j])
		}
		bottomRange--

		for i = bottomRange; i >= topRange; i-- {
			result = append(result, matrix[i][leftRange])
		}
		leftRange++

	}
	return result[:len(matrix[0])*len(matrix)]
}

// 剑指 Offer 32 - I. 从上到下打印二叉树
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	a := []TreeNode{}
	a = append(a, *root)
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

// 剑指 Offer 32 - II. 从上到下打印二叉树 II
//  关键点：如何确定内层的队列，就是该层全部元素，本次的循环范围，是上次的入队，上次的入队，是本次的层
func levelOrder2(root *TreeNode) [][]int {
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

func levelOrder3(root *TreeNode) [][]int {
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

			if level%2 == 0 {
				// 取值
				results[level] = append(results[level], v.Val)
			} else {
				results[level] = append([]int{v.Val}, results[level]...)
			}

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

// 剑指 Offer 33. 二叉搜索树的后序遍历序列
// 搜索树的后续满足 ：小-大-中 的顺序  递归
func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 || len(postorder) == 1 || len(postorder) == 2 || len(postorder) == 3 {
		return true
	}
	max := postorder[len(postorder)-1]
	split := len(postorder) - 1
	for i := 0; i < len(postorder)-1; i++ {
		if postorder[i] > max {
			split = i
			for j := split + 1; j < len(postorder)-1; j++ {
				if postorder[j] < max {
					return false
				}
			}
			break
		}
	}

	return verifyPostorder(postorder[:split]) && verifyPostorder(postorder[split:len(postorder)-1])

}

// 剑指 Offer 34. 二叉树中和为某一值的路径
// 关键点：看清题目，一定要求遍历到最后的叶节点;节点可能为负，所以没办法提前判断sum满足度而进行剪枝操作
func pathSum(root *TreeNode, sum int) [][]int {

	nums := []int{}
	results := &[][]int{}
	pathSumRecur(root, sum, nums, results)

	return *results
}

func pathSumRecur(root *TreeNode, sum int, nums []int, results *[][]int) {
	if root == nil {
		return
	}
	if sum == root.Val && root.Left == nil && root.Right == nil {

		nums = append(nums, root.Val)
		var temp = make([]int, len(nums))
		copy(temp, nums)
		*results = append(*results, temp)
		return
	} else {
		nums = append(nums, root.Val)

		pathSumRecur(root.Left, sum-root.Val, nums, results)
		pathSumRecur(root.Right, sum-root.Val, nums, results)
		return
	}

}

// 剑指 Offer 35. 复杂链表的复制
// 思路：先考虑简单链表的复制
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}

	// result := Node{}
	guard := Node{}
	guard.Next = head
	pointer := head
	prev := guard
	for pointer != nil {
		temp := Node{}
		temp.Val = pointer.Val
		*prev.Next = temp
		pointer = pointer.Next
	}

	return guard.Next

}

// 剑指 Offer 39. 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	m := map[int]int{}

	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}

	return -1

}

// 剑指 Offer 40. 最小的k个数
// 不调用系统函数，则考虑用快排、堆来求解
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func getLeastNumbers(arr []int, k int) []int {

	if false {
		sort.Ints(arr)
		return arr[:k]
	}

	result := []int{}
	if true {
		// result := []int{}
		in := &IntHeap{}
		for _, a := range arr {
			*in = append(*in, a)
		}

		heap.Init(in)
		for i := 0; i < k; i++ {
			result = append(result, heap.Pop(in).(int))
		}
	}
	return result
}

// 剑指 Offer 42. 连续子数组的最大和
// 关键点 找到状态转移表达式
func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			temp := nums[i]
			nums[i] = nums[i-1] + temp
		} else {
			temp := nums[i]
			nums[i] = temp
		}
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

// 剑指 Offer 50. 第一个只出现一次的字符
func firstUniqChar(s string) byte {

	if s == "" {
		return byte(' ')
	}

	type AAA struct {
		count int
		index int
	}

	m := map[byte]AAA{}

	for i, c := range s {
		temp := AAA{count: m[byte(c)].count + 1, index: i}
		m[byte(c)] = temp
	}

	var first byte = ' '
	firstIndex := len(s)
	for k, v := range m {
		if v.count == 1 {
			if v.index < firstIndex {
				firstIndex = v.index
				first = k
			}
		}
	}
	return first

}

// 剑指 Offer 52. 两个链表的第一个公共节点
// 关键点：链表第一个公共节点之后，都是相同的，所以要想办法拿到链表长度之差之后，再处理，比官方的所谓双指针好理解很多
// 或者直接想到用map记录一个链表的数据，第二个链表去查哈希，空间复杂度高一些
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	lenA := 0
	lenB := 0
	AA := headA
	BB := headB
	for headA != nil {
		lenA++
		headA = headA.Next

	}
	headA = AA

	for headB != nil {
		lenB++
		headB = headB.Next
	}
	headB = BB

	if lenA < lenB {
		for i := 0; i < lenB-lenA; i++ {
			headB = headB.Next
		}
	} else {
		for i := 0; i < lenA-lenB; i++ {
			headA = headA.Next
		}
	}

	for headA != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil

}

// 剑指 Offer 53 - I. 在排序数组中查找数字 I
func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] == target {
		return 1
	}
	count := 0
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			// 前后看
			count++
			for i := mid - 1; i >= 0 && nums[i] == nums[mid]; i-- {
				count++
			}
			for i := mid + 1; i < len(nums) && nums[i] == nums[mid]; i++ {
				count++
			}
			break
		}
	}
	return count

}

func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == mid {
			left = mid + 1
		} else if nums[mid] > mid {
			right = mid - 1
		} else {

		}
	}
	return left
}

// 剑指 Offer 54. 二叉搜索树的第k大节点
// 思路：中序遍历可以做到从小到大，从而找到第K大
func kthLargest(root *TreeNode, k int) int {
	resultArray := []int{}

	kthLargestRecur(root, &resultArray)

	return resultArray[len(resultArray)-k]
}
func kthLargestRecur(root *TreeNode, a *[]int) {

	if root == nil {
		return
	}

	kthLargestRecur(root.Left, a)

	*a = append(*a, root.Val)

	kthLargestRecur(root.Right, a)

}

// 剑指 Offer 55 - I. 二叉树的深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

// 剑指 Offer 55 - II. 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, isBalanced := isBalancedRecur(root)
	return isBalanced
}
func isBalancedRecur(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, isLeftBalanced := isBalancedRecur(root.Left)
	rightDepth, isRightBalanced := isBalancedRecur(root.Right)

	if !isLeftBalanced || !isRightBalanced {
		return -1, false
	}

	if math.Abs(float64(leftDepth-rightDepth)) <= 1 {
		return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1, true
	} else {
		return -1, false
	}

}

// 剑指 Offer 57. 和为s的两个数字
func twoSum(nums []int, target int) []int {

	left := 0
	right := len(nums) - 1

	for left < right {
		if target == nums[right]+nums[left] {
			break
		}

		if nums[right] >= target {
			right--
		} else {
			if target < nums[left]+nums[right] {
				right--
			} else {
				left++
			}
		}
	}
	return []int{nums[left], nums[right]}

}

// 剑指 Offer 57 - II. 和为s的连续正数序列
// 估计有数学公式可以弄，但我们找通用的计算机处理方法。
func findContinuousSequence(target int) [][]int {
	result := [][]int{}
	left := 1
	right := 2
	for left < right && right < (target+1)/2+1 {
		if (left+right)*(right-left+1)/2 == target {
			temp := []int{}
			for i := left; i < right+1; i++ {
				temp = append(temp, i)
			}
			result = append(result, temp)
			left++
			right = left + 1

		} else if (left+right)*(right-left+1)/2 > target {
			left++
			right = left + 1
		} else {
			right++
		}

	}
	return result
}

// 剑指 Offer 58 - I. 翻转单词顺序
func reverseWords(s string) string {
	sTrim := strings.Trim(s, " ")
	if sTrim == "" {
		return ""
	}
	result := ""
	stack := []string{}
	tempString := []byte{}
	flag := 0
	for i := 0; i < len(sTrim); i++ {
		if sTrim[i] == ' ' && flag == 0 {
			stack = append(stack, string(tempString))
			tempString = []byte{}
			flag = 1
		} else if sTrim[i] == ' ' && flag == 1 {
			continue
		} else {
			tempString = append(tempString, sTrim[i])
			flag = 0
		}

		if i == len(sTrim)-1 {
			// tempString = append(tempString, sTrim[i])
			stack = append(stack, string(tempString))

		}
	}
	for i := len(stack) - 1; i >= 0; i-- {
		result = result + stack[i] + " "
	}
	return result[:len(result)-1]

}

// 剑指 Offer 58 - II. 左旋转字符串
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// 剑指 Offer 60. n个骰子的点数
func twoSum111(n int) []float64 {
	if n == 1 {
		result := []float64{}
		for i := 0; i < 6; i++ {
			result = append(result, 1.0/6)
		}
		return result
	}
	result := []float64{}
	m := map[int]float64{}
	temp := twoSum111(n - 1)
	for i := 0; i < len(temp); i++ {
		for j := 0; j < 6; j++ {
			m[i+j+n] = m[i+j+n] + 1.0/6*temp[i]
		}
	}
	for i := 0; i < len(m); i++ {
		result = append(result, m[i+n])
	}

	return result
}

// 剑指 Offer 61. 扑克牌中的顺子
func isStraight(nums []int) bool {
	sort.Ints(nums)
	bigboom := 0
	if nums[0] == 0 {
		bigboom++
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			bigboom++
		} else {
			if nums[i-1] == 0 || nums[i] == nums[i-1]+1 {
				continue
			} else {
				if nums[i]-nums[i-1] > 0 && nums[i]-nums[i-1] <= bigboom+1 {
					bigboom -= nums[i] - nums[i-1] + 1
					continue
				} else {
					return false
				}

			}
		}
	}
	return true
}

// 剑指 Offer 62. 圆圈中最后剩下的数字
// 关键点：画图试图举例推导出递归表达式
func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	return (lastRemaining(n-1, m) + m) % n

}

// 剑指 Offer 65. 不用加减乘除做加法
// 关键点： 不能直接加，只能递归复用temp和carry的方法，实现temp和caryy的相加，直到carry为0
func add(a int, b int) int {
	// var x int = 0xFFFFFFFF
	// carry := -1

	// 1. 按位和加上进位和，但是进位和用到了+
	// tempSum := a ^ b
	// carry := a & b << 1
	// result := tempSum+carry

	// 2. 把加法递归复用现在的实现。
	// if b == 0 {
	// 	return a
	// }
	// tempSum := a ^ b
	// carry := a & b << 1
	// return add(tempSum, carry)

	// 3.考虑把递归转换成循环迭代实现
	for b != 0 {
		tempSum := a ^ b
		carry := a & b << 1
		a = tempSum
		b = carry
	}
	return a

}

// 剑指 Offer 66. 构建乘积数组
// 思路：有重复计算，肯定跟动态规划有关
func constructArr(a []int) []int {

	return []int{}
}

// 235. 二叉搜索树的最近公共祖先
// 剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}

}

// 236. 二叉树的最近公共祖先
// 剑指 Offer 68 - II. 二叉树的最近公共祖先
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil && right == nil {
		return left
	} else {
		return right
	}

}

// 剑指 Offer 14- I. 剪绳子
func cuttingRope(n int) int {
	//  2 1
	//  3,2
	//  4,4
	//  5,6   2*3
	//  6,9   3*3
	//  7,12  3*4
	//  8,16  4*4
	//  9,27  3*3*3
	//  10,36 3*4*3
	// 有效拆分中的最大的那个数：cuttingRope(n-1)
	if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	} else if n == 4 {
		return 4
	}
	max1 := math.Max(float64(cuttingRope(n-1)+n-1), float64(cuttingRope(n-2)+2*(n-2)))
	max2 := math.Max(float64(cuttingRope(n-3)+3*(n-3)), max1)
	return int(max2)

}

// 剑指 Offer 64. 求1+2+…+n
func sumNums(n int) int {

	_ = n > 0 && func() bool {
		n = sumNums(n-1) + n
		return true
	}()
	return n
}

// 剑指 Offer 16. 数值的整数次方
func myPow(x float64, n int) float64 {
	result := 1.0
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n > 0 {
		temp := myPow(x, n/2)
		result = temp * temp * myPow(x, n%2)
	}
	if n < 0 {
		m := -n
		temp := myPow(x, m/2)
		result = 1.0 / (temp * temp * myPow(x, m%2))

	}
	return result
}

// 剑指 Offer 20. 表示数值的字符串
func isNumber(s string) bool {

	match, _ := regexp.Match("^[+-]?(([0-9]+)|([0-9]+\\.[0-9]*)|([0-9]*\\.[0-9]+))([eE][+-]?[0-9]+)?$", []byte(s))
	return match

}

// 剑指 Offer 26. 树的子结构
// 关键点：递归中嵌套递归
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	}
	if A == nil || B == nil {
		return false
	}

	result := false
	if A.Val == B.Val {
		result = helper(A, B)

	}
	if !result {
		return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
	}

	return result

}

//helper 校验 B 是否与 A 的一个子树拥有相同的结构和节点值
func helper(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	//a.Val == b.Val 递归校验 A B 左子树和右子树的结构和节点是否相同
	return helper(a.Left, b.Left) && helper(a.Right, b.Right)
}

// 剑指 Offer 31. 栈的压入、弹出序列
func validateStackSequences(pushed []int, popped []int) bool {
	temp := []int{}
	popIndex := 0
	for _, v := range pushed {
		temp = append(temp, v)

		for popIndex < len(popped) && len(temp) > 0 && popped[popIndex] == temp[len(temp)-1] {
			popIndex++
			temp = temp[:len(temp)-1]
		}
	}
	if len(temp) == 0 && popIndex == len(popped) {
		return true
	} else {
		return false
	}
}

// 剑指 Offer 38. 字符串的排列
// func permutation(s string) []string {
// 	result := []string{}

// 	for i := 0; i < len(s); i++ {

// 	}

// }
