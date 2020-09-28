package main

import (
	"fmt"
	"os"
)

const (
	Monday = 1 << iota
	Tuesday
	Wednesday
)

func main1() {
	fmt.Println(os.Args)
	a, b := "a", "c"
	a, b = b, a
	fmt.Println("hello world")
	fmt.Println(a, b, Tuesday, Wednesday)

	//基本语言类型
	// bool a = true
	// string b = "ccc"
	// int c = 12
	// fmt.Println(a,b,c)

}
