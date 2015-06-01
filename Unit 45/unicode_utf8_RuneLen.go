package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s string = "한"
	fmt.Println(len(s)) // 3: 한글은 3바이트로 저장하므로 3

	var r rune = '한'
	fmt.Println(utf8.RuneLen(r)) // 3: 한글은 3바이트로 저장하므로 3
}
