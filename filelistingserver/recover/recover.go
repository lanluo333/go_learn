package main

import (
	"errors"
	"fmt"
)

func tryRecover(){

	defer func() {
		// recover仅在defer调用中使用
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		}else {
			panic(r)
		}
	}()

	/**
	panic会停止当前函数的执行
	然后一直向上返回，执行每一层的defer
	如果没有遇见recover，程序退出
	 */
	panic(errors.New("this is error"))
}

func main(){
	tryRecover()
}

