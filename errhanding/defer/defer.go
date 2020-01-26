package main

import (
	"bufio"
	"fmt"
	"learn/functional/fib"
	"os"
)

func tryDefer(){
	// defer是等函数结束的时候才调用的
	// 而defer相当于一个栈，先进后出，所以会按321的顺序输出
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)

	for i:=0;i<100 ;i++  {
		defer fmt.Println(i)
		if i == 30 {
			panic("too many")
		}
	}

}


func writeFile(filename string){
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//panic(err)
		//fmt.Println("Error:",err.Error())
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		}else {
			fmt.Printf("%s, %s, %s\n",pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}

	defer file.Close()

	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib.Fibonacci()
	for i :=0; i< 20 ; i++  {
		fmt.Fprintln(write, f())
	}

}

func main()  {
	//tryDefer()
	writeFile("fib.txt")
}

