package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 10进制转换为二进制
func converToBin(n int) string {
	result := ""
	for ; n>0 ; n/=2  {
		lsb := n % 2
		// strconv.Itoa是把整数换成字符串
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func forever(){
	// 死循环
	for {
		fmt.Println("forever")
	}
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err!= nil {
		panic(err)
	}

	printFileContents(file)
	/*
	scanner := bufio.NewScanner(file)

	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}*/
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
			converToBin(5),  // ☆
			converToBin(13),
		)
	printFile("learn/basc/abc.txt")

	fmt.Println()

	s := `abc"d"
	kkk
	123
	p`
	printFileContents(strings.NewReader(s))
}


