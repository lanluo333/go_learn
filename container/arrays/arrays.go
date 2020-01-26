package main

import "fmt"

// [5]表示的是数组 []表示的是splice
func printArr(arr [5] int){
	for i, v:=range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	// 使用:=来声明数组的话，要在后面先赋值
	arr2 := [3]int{1,3,5}
	// ...就是让编译器来帮我们数有几个值
	arr3 := [...]int{2, 4, 6, 8, 10}

	// 声明4行5列的一个数组
	var grid [4][5] int

	fmt.Println(arr1, arr2, arr3)  // [0 0 0 0 0] [1 3 5] [2 4 6 8 10]
	fmt.Println(grid)  // [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]

	// 遍历数组
	for i:=0; i<len(arr2); i++ {
		fmt.Println(arr2[i])
	}
	// 以上输出如下
	// 1
	// 3
	// 5

	fmt.Println()
	// 另一种遍历方式
	for i, v:=range arr2 {
		fmt.Println(i, v)
	}
	// 以上输出
	// 0 1
	// 1 3
	// 2 5

	fmt.Println()
	// 遍历的时候不需要下标，只取值
	for _, v:=range arr2 {
		fmt.Println(v)
	}
	// 以上输出如下
	// 1
	// 3
	// 5

	println("================================")
	printArr(arr1)
	// 以上输出
	// 0 0
	// 1 0
	// 2 0
	// 3 0
	// 4 0

	fmt.Println()
	printArr(arr3)
	// 0 2
	// 1 4
	// 2 6
	// 3 8
	// 4 10

}
