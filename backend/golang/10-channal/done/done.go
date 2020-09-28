package main

import (
	"fmt"
	"sync"
)

func worker(id int,c chan int,wg *sync.WaitGroup) {
	for cin := range c { // range去检测channel 是否close
		//rcv,ok := <-cin

		fmt.Printf("id:%d,rcv:%d \n",id,cin)
		wg.Done()
	}
}
type workerx struct {
	in chan int
	wg *sync.WaitGroup
}
func createWorker(id int,wg *sync.WaitGroup) workerx {
	w:= workerx{
		in:make(chan int),
		wg: wg,
	}
	go worker(id,w.in,w.wg)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]workerx
	//var channels [10] chan int
	for i := 0; i < 10; i++ {
		workers[i]=createWorker(i,&wg)
	}


	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- i+'a'

	}
	for i := 0; i < 10; i++ {
		workers[i].in <- i+'A'

	}
	wg.Wait()



}

func main() {
	chanDemo()
	//bufferedChannel()
}