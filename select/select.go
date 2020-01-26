package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int){
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

func generator() chan int{
	out := make(chan int)
	go func() {
		i := 0
		for  {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main()  {
	var c1, c2  = generator(), generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		// 哪个数据先被接收就选择哪个
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			// 800毫秒没有生产数据就打印timeout
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len = ",len(values))
		case <-tm:
			fmt.Println("Bye")
			return
		}
	}
}


