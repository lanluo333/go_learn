package main

import (
	"fmt"
	"time"
)

/**
 * 通道
 */

func worker(id int, c chan int){
	/*
	for  {
		// 找个人来专门负责接收
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("worker %d received %c \n", id, n)
	}*/

	for n := range c{
		fmt.Printf("worker %d received %c \n", id, n)
	}

}

// chan<-表示返回的chan只能往chan里面存数据，不能从里面取数据
func createWorker(id int) chan<- int{
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo(){
	var channels [10]chan<- int
	for i := 0;i<10 ;i++  {
		/*
		channels[i] = make(chan int)
		go worker(i, channels[i])*/
		channels[i] = createWorker(i)
	}

	for i:=0; i<10; i++  {
		channels[i] <- 'a' + i
	}

	for i:=0; i<10; i++  {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel(){
	// 3表示的是缓冲，表示在没有人接收的情况下，可以往chan里面连续放3个数据而不取出
	c := make(chan int, 3)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func closeChannel(){
	// 3表示的是缓冲，表示在没有人接收的情况下，可以往chan里面连续放3个数据而不取出
	c := make(chan int, 3)
	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main()  {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
	// worker 0 received a
	//worker 1 received b
	//worker 2 received c
	//worker 3 received d
	//worker 3 received D
	//worker 4 received e
	//worker 4 received E
	//worker 5 received f
	//worker 5 received F
	//worker 6 received g
	//worker 6 received G
	//worker 7 received h
	//worker 7 received H
	//worker 1 received B
	//worker 0 received A
	//worker 9 received j
	//worker 2 received C
	//worker 8 received i
	//worker 8 received I
	//worker 9 received J

	fmt.Println("========================")

	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("========================")
	// Buffered channel
	//worker 0 received a
	//worker 0 received b
	//worker 0 received c
	//worker 0 received d

	fmt.Println("Channel close and range")
	closeChannel()

	// Channel close and range
	//worker 0 received a
	//worker 0 received b
	//worker 0 received c
	//worker 0 received d

}

