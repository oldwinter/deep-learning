package main

import (
	"fmt"
	"time"
)

func main() {
	a := [1000]int{}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				a[i]++
				fmt.Printf("hello from:%d\n",i)
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(10*time.Minute)
	fmt.Println(a)
}