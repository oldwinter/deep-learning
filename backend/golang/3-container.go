package main

import (
	"fmt"
)

func main1() {

	// 3-1 数组
	var arr1 [5]int
	arr2 := [2]int{1, 2}
	arr3 := [...]int{1, 3, 4}
	var grid [3][4][5]bool
	fmt.Println(arr1, arr2, arr3, grid)
	sum := 0
	for _, v := range arr3 {
		sum += v
	}
	// !数组会拷贝，除非换成*
	printArr := func(arr [3]int) int {
		return arr[0]
	}
	printArr(arr3)

	// 3-2 切片 核心：切片是数组的一个view，是ptr+len+cap组成的数据结构
	arr := [...]int{0, 1, 2, 3}
	slice := arr[1:2]
	fmt.Println(slice)
	// 数组变切片
	slice11 := arr2[:]
	fmt.Println(slice11)
	var s []int //s == nil
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	slice2 := make([]int, 10, 32)
	fmt.Println(slice2)
	copy(slice2, slice11)
	slice2 = append(slice2[1:2], slice11...)
	// 3-3 切片删除某元素、模拟队列和堆栈
	s33 := []int{0, 1, 2, 3, 4, 5}
	s33 = append(s33[:3], s33[4:]...)
	front := s33[0]
	s33 = s33[1:]
	end := s33[len(s33)-1]
	s33 = s33[:len(s33)-1]
	fmt.Println(front, end)

	// 3-4 Map
	m := map[string]string{
		"name": "cdd",
	}
	m2 := make(map[string]string)
	var m3 map[string]string

	value, isExist := m["value"]
	fmt.Println(m, m2, m3, ";", m["name"], m["noPresent"], value, isExist)
	delete(m, "name")
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 3-5 Map寻找最长不含有重复字符的子串 例题
	// 3-6 字符串处理
	ss35 := "rugo 有明天~！"
	for i, v := range []rune(ss35) {
		fmt.Printf("%d %x ", i, v)
	}

	// 5-1 接口

	type Retriver interface {
		Get(URL string) string
	}
	download = func(r *Retriver) string {
		return r.Get("www.cdd.com")
	}

	type RetriverNew type {
		s string
	}
	func (r *Retriver)Get() {
		fmt.Println(r)
	}
	var r Retriver
	r = RetriverNew.Get()
}
