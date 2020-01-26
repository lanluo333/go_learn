package main

/**
 * 字符串
 */

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱!"
	fmt.Println(s)  // 输出 "Yes我爱!"

	for i, ch := range s {
		// ch is rune
		// %x 每个字节用两字符十六进制数表示（使用a-f）
		fmt.Printf("(%d, %x)", i, ch)
		// 输出 (0, 59)(1, 65)(2, 73)(3, 6211)(6, 7231)(9, 21)
	}

	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 输出 Rune count: 6
	fmt.Println()

	// 直接把字符串转成rune类型，然后就可以支持中文了
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	// 输出 (0 Y) (1 e) (2 s) (3 我) (4 爱) (5 !)

}
