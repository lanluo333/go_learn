package queue

import "fmt"

func ExampleQueue_pop(){
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())

	// OutPut:
	// 1
	// 2
	// false
	// 3
}
