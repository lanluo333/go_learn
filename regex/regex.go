package main

import (
	"fmt"
	"regexp"
)

const text  = "My email is 1234567@qq.com"

func main()  {
	re := regexp.MustCompile("@qq.com")
	match := re.FindString(text)
	fmt.Println(match)
}

