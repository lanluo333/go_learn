package main

import (
	"fmt"
	"runtime"
	"time"
)

func main()  {
	var a [10]int
	for i:=0; i<10; i++ {
		go func(i int) {
			for  {
				//fmt.Printf("Hello from goroutine %d \n", i)
				a[i]++
				// 协程交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)  // [411 388 455 446 388 393 362 366 396 405]
}

