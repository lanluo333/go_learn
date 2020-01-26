package main

import (
	"fmt"
	"sync"
)

/**
不要通过共享内存来通信，要通过通信来共享内存
 */

func doWorker(id int, c chan int, w worker){

	for n := range c{
		fmt.Printf("worker %d received %c \n", id, n)
		w.done()
		/*
		go func() {
			done <- true
		}()*/
	}

}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker{
	w := worker{
		in : make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w.in, w)
	return w
}

func chanDemo(){
	var workers [10] worker
	var wg sync.WaitGroup
	for i := 0;i<10 ;i++  {
		/*
		channels[i] = make(chan int)
		go worker(i, channels[i])*/
		workers[i] = createWorker(i, &wg)
	}

	// 添加20个任务
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers  {
		worker.in <- 'A' + i
	}
	wg.Wait()

	// 等所有都结束以后在并发发出去
	/*
	for _, worker := range workers {
		// 因为大写的A和a总共写进去了两遍，所以这里也要读出来两次
		<-worker.done
		<-worker.done
	}*/
}

func main()  {
	chanDemo()
}

