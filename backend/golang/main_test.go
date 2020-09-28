package main

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestAbc(t *testing.T) {
	t.Log("aaaa")
}

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int
	b = a
	t.Log(a, b)

	var c int64
	b = int(c)
	t.Log(b)

	t.Log(math.MaxFloat64)

}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "8")
	t.Log(len(s))
}

func TestComppareArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{4, 5, 2}
	t.Log(a, b)

}

func TestComppareArray111(t *testing.T) {
	a := 7 //0111
	a = a & (1 << 2)
	t.Log(a)

}

func TestXunHuawei(t *testing.T) {

	if true {
		t.Log("daf")
	}
	if c := "c"; c == "a" {
		t.Log("yes")
	}
	for i := 0; i < 11; i++ {
		t.Log("ff")
	}
	a := 2
	switch a {
	case 1, 2:
		t.Log("1")
	case 3, 5:
		t.Log("2")
	}
}

func TestShuzu(t *testing.T) {
	var a [5]int
	a[0] = 1
	b := [...]int{0, 1, 2, 3}

	c := [2][2]int{{1}, {2}}
	t.Log(a, b, c)

	for idx, xx := range b {
		t.Log(idx, xx)
	}

	arr := [5]int{1, 2, 3, 4, 5}
	arr_aa := arr[:5]
	t.Log(arr_aa)

}

func TestSlice(t *testing.T) {
	var s0 []int
	s0 = append(s0, 2)
	t.Log(s0, len(s0), cap(s0))
	s2 := make([]string, 0, 10)
	s2 = append(s2, "adsfa", "fsfds")
	s2 = append(s2, "a")
	t.Log(s2, len(s2), cap(s2))

	// for i := 0; i < 100; i++ {
	// 	s2 = append(s2, string(i))
	// 	t.Log(len(s2), cap(s2))
	// }

	months := [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	a := months[2:7]
	b := months[4:8]
	t.Log(a, len(a), cap(a))
	t.Log(b, len(b), cap(b))
	months[5] = 1
	t.Log(a, b)

}

func TestMap(t *testing.T) {
	// m := map[int]int{}
	m := map[string]int{"a": 1, "b": 2, "asfds": 3}
	v, ok := m["a"]
	v2, ok2 := m["b"]
	t.Log(v, ok, v2, ok2)
	t.Log(len(m), m["b"])
	for k, v := range m {
		t.Log(k, v)
	}
}

func TestFactory(t *testing.T) {
	m := map[int]func(in int) int{}
	m[1] = func(in int) int { return in }
	m[2] = func(in int) int { return in * in }
	m[3] = func(in int) int { return in * in * in }
	t.Log(m[1](2), m[2](3))
}

func TestMapForSet(t *testing.T) {
	sets := map[int]bool{}
	sets[1] = true
	sets[2] = true
	delete(sets, 2)
	t.Log(sets)
}

func TestStringa(t *testing.T) {
	// var s string
	// t.Log(s)
	// s = "hello"
	// t.Log(len(s))
	// s = "中国"
	// t.Logf("%x,%x,%x", s[0], s[1], s[2])

	// c := []rune(s)
	// t.Logf("%x %x", c[0], c[1])

	// s1 := "a,b,c"
	// p := strings.Split(s1, ",")
	// for _, v := range p {
	// 	t.Log(v)
	// }

	// s = "我欲乘风破浪"
	// pp := strings.Split(s, "欲")
	// for _, c := range pp[] {
	// 	t.Logf("%[1]c %[1]x", c)
	// }

	aa := strconv.Itoa(10)
	t.Log(aa)

	if _, err := strconv.Atoi("10000000000000"); err != nil {
		t.Log(err)
	}
	// t.Log(err)
}
func returntwoss(a int) int {
	return (10)

}

func funcinfunc(infun func(a int) int) func(a int) int {
	return func(n int) int {
		infun(10)
		return 100
	}

}

func TestFuncaa(t *testing.T) {
	a := funcinfunc(returntwoss)(20)
	t.Log(a)
	// func returntwoss() (int int)  {
	// 	return rand.Intn(10), rand.Intn(20)

	// }
	// t.Log(returntwoss())
}

func sum(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func aaattt(a int) int {
	fmt.Println(a)
	return a + 10
}

func TestMultiOpt(t *testing.T) {
	defer aaattt(1)
	t.Log("start")
	panic("fasf")

	// t.Log(sum(1, 2, 3, 4, 5))
}
