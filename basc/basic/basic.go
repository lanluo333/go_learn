package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 一次声明多个变量，避免一直使用var关键字
// 不过，go语言里的这些变量在外面声明的话，不是全局变量的意思，而是包变量，也就是这个包里的变量
var (
	aa = 33
	ss = "kk"
	bb = true
)

func variableZeroValue(){
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}

func euler(){
	// 这里是复数
	c := 3 + 4i
	println(cmplx.Abs(c))
}

func triangle(){
	var a, b int = 3, 4
	println(calcTriangle(a, b))
}

func calcTriangle(a int, b int) int {
	var c int
	// go语言中的sqrt里面的参数和返回的值都是float类型
	// go语言里面的类型转换没有隐式转换，只有强制转换
	c = int(math.Sqrt(float64(a * a + b * b)))
	return c
}

// 常量
func consts(){
	//const filename  = "abc.txt"
	// 如果常量后面没有声明具体是什么类型，则常量便是文本类型
	//const a, b  = 3, 4
	const(
		filename  = "abc.txt"
		a, b  = 3, 4
	)
	var c int
	// 由于常量a,b没有声明是什么类型，所以可以被自动替换成其他类型
	c = int(math.Sqrt(float64(a * a + b * b)))
	println(filename, c)
}

// 枚举
func enums(){
	/*
	const(
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)*/
	// iota会从0开始，然后逐渐的往下递增
	const(
		cpp = iota
		java
		python
		golang
	)
	// iota也可以用在表达式上面
	const(
		b = 1 << (10 * iota)  // 这里会逐渐的往左移1位、10位、20位等
		kb
		mb
		gb
		tb
		pb
	)
	println(cpp, java, python, golang)
	println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	println(aa, ss, bb)

	println("=============================")
	euler()

	triangle()

	println("=========================")
	consts()

	println("==========================")
	enums()
}

func variableTypeDeduction(){
	// 不规定类型
	var a, b, c, s = 3, 4, true, "string"
	fmt.Println(a, b, c, s)
}

func variableInitValue(){
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableShorter(){
	// 使用:=和使用var进行声明效果是一样的,如果是在函数外，则不能使用这种方式声明，必须要var关键字
	a, b, c, s := 3, 4, true, "string"
	b = 5
	fmt.Println(a, b, c, s)
}

