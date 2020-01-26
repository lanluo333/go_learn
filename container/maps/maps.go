package main

import "fmt"

func main() {
	// []里的string表示的是key的类型，外面的string表示的是value的类型
	m := map[string]string{
		"name" : "rui",
		"course" : "golang",
		"site" : "golang",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]string // m3 ==nil

	fmt.Println(m, m2, m3)  // map[course:golang name:rui site:golang] map[] map[]
	fmt.Println()

	// 遍历map
	fmt.Println("Traversing map")
	// map是无序的，所以打印出来的话不一定是按照原来的顺序那样输出
	// 如果需要有序，则需要手动排序，也就是把他排序放进slice里，然后再遍历slice
	for k,v := range m {
		fmt.Println(k, v)
	}
	// 以上输出
	// name rui
	// course golang
	// site golang
	fmt.Println()

	fmt.Println("Getting Value")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)  // golang true
	// 取map里面不存在的key,不会报错，只会拿出ZeroValue,也就是空值
	// 也就是说如果值不存在，第二个接收的参数(ok)会返回false
	caurseName, ok := m["caurse"]
	fmt.Println(caurseName, ok)  // 输出  false

	fmt.Println()

	// 以下输出 key not exist
	if caurseName, ok := m["caurse"]; ok {
		fmt.Println(caurseName, ok)
	}else {
		fmt.Println("key not exist")
	}

	fmt.Println()

	fmt.Println("Delete Value")
	name, ok := m["name"]
	fmt.Println(name, ok) // rui true

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok) //   false

	// map使用的是哈希表，所以必须比较相等才能取出值
	// 也就是说，除了map,slice,function的内建类型以外，其他类型都可以作为key
	// struct自建类型只要不包含以上类型也可以当做是key

}
