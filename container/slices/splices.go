package main

import "fmt"

/**
数组切片
 */

// []代表的是slice类型,slice是array的视图
func updateSlice(s []int){
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6]) // 输出 arr[2:6] = [2 3 4 5]
	fmt.Println("arr[:6] =", arr[:6])  // 输出 arr[:6] = [0 1 2 3 4 5]
	fmt.Println("==========================")

	s1 := arr[2:]
	fmt.Println("s1 =", s1)  // 输出 s1 = [2 3 4 5 6 7]

	s2 := arr[:]
	fmt.Println("s2 =", s2)  // s2 = [0 1 2 3 4 5 6 7]

	fmt.Println("updateSlice(s1)")
	// slice是array的底层视图，相当于在slice里面改变值之和，原本的array也会改变
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2) // [100 3 4 5 6 7]
	fmt.Println(arr) // [0 1 100 3 4 5 6 7]

	fmt.Println("SliceExtend")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	fmt.Println(s1)  // [100 1 100 3 4 5 6 7]
	// s2的值会是[5 6],虽然s1里面没有6的值，但其实slice是可以向后扩展的。
	// 也就是说slice还是可以取到原来的数组的后面的一些值，但s1[4]这种形式就不行，s1只能看到自己取到的值
	// 尽管slice可以向后扩展，但不可以超过原来底层数组的最大长度
	// 即s[1]不可以超越len(s),向后扩展不可以超越底层数组cap(s)
	s2 = s1[3:5]
	fmt.Println(s2) // [5 6]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1)) // s1=[2 3 4 5], len(s1)=4, cap(s1)=6
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2)) // s2=[5 6], len(s2)=2, cap(s2)=3

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5=",s3, s4, s5)  // s3,s4,s5= [5 6 10] [5 6 10 11] [5 6 10 11 12]
	// s4, s5已经和arr不同了,s3添加元素时由于再添加一个元素不会超越cap，所以，原来的arr最后的7会被覆盖成10
	// 添加元素时如果超越cap，系统会重新分配更大的数组
	fmt.Println("arr=", arr)  // arr= [0 1 2 3 4 5 6 10]
}
