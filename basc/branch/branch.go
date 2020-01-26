package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		// panic的作用就是中断程序，然后报错
		panic(fmt.Sprintf("Wrong score: %d",score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	path, _ := os.Getwd()
	filename  := path + "/basc/abc.txt"

	contents,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}

	// 也可以如下写法,但因为下面的contents是在if那里定义的，所以其作用于只在那个if的作用域里
	/*
	if contents,err := ioutil.ReadFile(filename);err != nil{
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}*/
	fmt.Println(
			grade(59),
			grade(60),
			grade(82),
			grade(90),
			grade(100),
			grade(101),
		)
}
