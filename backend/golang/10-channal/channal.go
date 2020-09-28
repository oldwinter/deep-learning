package main

import (
	"fmt"
	"time"
)

func worker(id int,c chan int) {
	for cin := range c { // range去检测channel 是否close
		//rcv,ok := <-cin

		fmt.Printf("id:%d,rcv:%d \n",id,cin)
	}
}

func chanDemo() {

	var channels [10] chan int
	for i := 0; i < 10; i++ {
		//c:= make(chan int)
		channels[i] = make(chan int)
		go worker(i,channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i
	}

	//n := <-c
	time.Sleep(time.Millisecond)
	//fmt.Println()
}

func bufferedChannel() {
	c:= make(chan int,3)
	go worker(0,c)
	c <- 1
	c <- 2
	c <- 3
	close(c) // 发送方才能close
	//c <- 4
	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	bufferedChannel()
}