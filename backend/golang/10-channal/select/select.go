package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate() chan int {
	out := make(chan int)
	go func() {
		i :=0
		for  {
			time.Sleep(time.Duration(rand.Intn(100000)))
			out <- i
			i++
		}
	}()
	return out
}



func main() {

	var c1,c2 = generate(),generate()
	//w:=

	for  {
		select {
		case n := <-c1:
			fmt.Println("c1",n)
		case n := <-c2:
			fmt.Println("c2",n)
		//default:
			//fmt.Println("nodata")
		}
	}



}
