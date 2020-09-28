package main

import "fmt"

type treeNode struct {
	Val         int
	Left, Right *treeNode
}

type point struct {
	i int
	j int
}

func (n *treeNode) print() {
	fmt.Println(n.Val)
}
func (n *treeNode) setVal(val int) {
	n.Val = val
}

func createNode(value int) *treeNode {
	return &treeNode{Val: value}
}

func main() {
	// 4-1结构体和方法
	var root treeNode
	root = treeNode{3, nil, nil}
	root.Left = &treeNode{}
	root.Right = &treeNode{5, nil, nil}
	root.Right.Left = new(treeNode)

	root.Right.Right = createNode(2)
	root.print()
	root.setVal(5)
	root.print()

	// 4-2包和封装
	// 4-3 扩展已有类型 组合 、别名
	type myTreeNode struct {
		my *treeNode
	}

	type Queue []int
	type myTreeNodeNew treeNode

}
