package main

import (
	"fmt"
	"testing"
)

type FuncAA func(a *People)

type Programmer interface {
	WriteHelloWorld() string
	GoHello() string
}

type GoProgrammer struct {
	Programmer
	GoNum int
}

func (gog *GoProgrammer) WriteHelloWorld() string {
	return "fmt.println\"aa\""
}

func (gog *GoProgrammer) GoHello() string {
	return "fmt.println\"aa\""
}

type JavaProgrammer struct {
	Programmer
	JavaNum int
}

func (gog *JavaProgrammer) WriteHelloWorld() string {
	return "javaprint"
}

func gogo(pp Programmer) {
	fmt.Println(pp.WriteHelloWorld())
}
func TestClient(t *testing.T) {
	var p, p2 Programmer
	p = new(GoProgrammer)
	p2 = new(JavaProgrammer)
	gogo(p)
	p.GoHello()
	gogo(p2)
	p2.GoHello()
	i, ok := p.(Programmer)
	t.Log(i, ok)
}

type People struct {
	Aaa string
	Bbb int
	Ccc int
}

func aaaa(p *People) {
	fmt.Println(p.Ccc)
}
