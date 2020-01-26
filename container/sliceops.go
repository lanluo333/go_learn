package container

import "fmt"

func printSlice(s [] int){
	fmt.Printf("%v,len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	fmt.Println("Create Slice")
	// 直接创建slice
	var s [] int  // Zero Value For slice is nil

	for i := 0; i<100; i++ {
		// 从打印可以看到，当slice装不下的时候，cap会以乘2的倍数向后扩展
		printSlice(s)
		s = append(s, 2 * i + 1)
	}
	fmt.Println(s)

	// 也可以使用以下的方法创建slice
	s1 := []int{2, 4, 6, 8} // 声明一个数组，然后把这个数组review到slice里
	printSlice(s1)

	// 直接使用make来进行slice，其中16是表示长度
	s2 := make([]int, 16)
	printSlice(s2)

	// make的第三个参数是cap的长度
	s3 := make([]int, 10, 32)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	// 把s2下标为3的元素删掉，然后把后面的挪到前面
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	// 取出slice的头尾
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	back := s2[len(s2) - 1]
	s2 = s2[:len(s2) - 1]
	fmt.Println(back)
	printSlice(s2)
}
