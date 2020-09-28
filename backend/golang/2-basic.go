package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/cmplx"
	"reflect"
	"runtime"
)

// 3. package 内变量，而不是全局变量 ; 无法使用:=定义
var bb = "cc"

const BB = "BB"

func main2() {
	// 2-1.变量定义
	fmt.Printf("%d %s\n", 1, "hello world")

	a, b := 1, "hello world"
	fmt.Println(a, b)

	// 2-2.内建变量类型
	var isExist bool = true
	var s string = "hello world"
	var a1 uint32 = 0
	var a2 int64 = 0xFFFFFFFF
	var a3 uintptr = 0xFFFFFFFF

	var b1 byte = 'b'
	var r rune = 'a' //实际底层用int32表示

	var complex complex128 = 4 + 2i
	result := cmplx.Pow(math.E, math.Pi*1i)

	fmt.Println(isExist, s, a1, a2, a3, b1, r, complex, result)

	var a6, b6 int64 = 3, 4
	var c int
	c = int(math.Sqrt(float64(a6*a6 + b6*b6)))
	fmt.Println(c)

	// 2-3.常量和枚举
	const f7 = "hello"
	const (
		AA = 'a'
		BB = 'b'
	)
	const (
		a8 = iota
		_
		b8
		c8
	)
	const (
		a88 = 1 << iota
		_
		b88
		c88
	)
	fmt.Println(a8, b8, c8, ";", a88, b88, c88)

	// 2-4 条件语句
	const filename = "basic.go"

	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("file:", string(content))
	}
	condition := 0
	switch {
	case condition < 0:
		fmt.Println(condition)
	case condition == 0:
		fmt.Println(condition)
	default:
		panic(condition)
	}

	// 2-5 循环语句
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	for {
		fmt.Println("a")
		break
	}

	// 2-6 函数
	func1 := func(a, b int) (int, int) {
		return a + b, a - b
	}
	fmt.Println(func1(1, 2))

	apply := func(op func(int, int) int, a, b int) int {
		p := reflect.ValueOf(op).Pointer()
		opName := runtime.FuncForPC(p).Name()
		fmt.Println(opName, a, b)
		return op(a, b)
	}
	apply(func(a, b int) int {
		return a + b
	}, 3, 4)

	// 2-7 指针
	swap := func(a, b *int) {
		*a, *b = *b, *a
	}
	a7, b7 := 1, 2
	swap(&a7, &b7)
	fmt.Println(a7, b7)
}
