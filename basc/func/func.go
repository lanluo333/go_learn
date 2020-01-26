package main

/**
* 函数
*/

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a,b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// 函数如果返回2个蚕食，有一个不想用的话就用_
		q, _ := div(a, b)
		return q, nil
	default:
		//panic("unsupported operation:"+op)
		return 0, fmt.Errorf("unsupported operation:%s", op)
	}
}

func div(a, b int) (int, int){
	return a / b, a % b
}

// 可以默认的给返回值取名
// 当然实际中如果函数里面代码量比较多得多，有时很难区分哪里给变量（q,r）赋值，所以还是建议以上的方法
func div_2(a, b int) (q, r int){
	q = a / b
	r = a % b
	return
}

func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("Calling %s with %d %d \n",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(), a, b)
	return op(a, b)
}

// ...表示的是可变参数列表，也就是可以输入任意多个参数进来
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

func swap(a, b *int){
	*b, *a = *a, *b
}

func main() {
	// 以下输出 unsupported operation:x
	if result, err := eval(3, 4, "x");err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}


	fmt.Println(div(13, 3))  // 输出 4 1
	q, r := div_2(13,3)
	fmt.Println(q, r)  // 输出 4 1

	fmt.Println(apply(func(a int, b int) int {
		 return int(math.Pow(float64(a), float64(b)))
	},3, 4))


	fmt.Println(sum(1, 2, 3, 4, 5))  // 输出15

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)  // 输出 4 3

}
