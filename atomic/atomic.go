package main

/*
 * 原子操作
 */

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock sync.Mutex // 互斥锁
}

func (a *atomicInt) increment(){
	// 加锁
	fmt.Println("safe increment")
	// 把defer操作控制在某个函数体里
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

func (a *atomicInt) get() int  {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main()  {
	var a atomicInt
	a.increment()  // 输出 safe increment
	go func() {
		a.increment()  // 输出 safe increment
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get()) // 输出 2
}

