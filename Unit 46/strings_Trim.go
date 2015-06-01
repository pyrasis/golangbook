package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Trim("Hello,~~ world!~~", "~~")) // Hello,~~ world!
	fmt.Println(strings.Trim("  Hello, world!  ", " "))  // Hello, world!

	fmt.Println(strings.TrimLeft("~~Hello,~~ world!~~", "~~"))  // Hello,~~ world!~~
	fmt.Println(strings.TrimRight("~~Hello,~~ world!~~", "~~")) // ~~Hello,~~ world!

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r) // r이 한글 유니코드이면 true를 리턴
	}
	var s = "안녕 Hello 고 Go 언어"
	fmt.Println(strings.TrimFunc(s, f))      //  Hello 고 Go
	fmt.Println(strings.TrimLeftFunc(s, f))  //  Hello 고 Go 언어
	fmt.Println(strings.TrimRightFunc(s, f)) // 안녕 Hello 고 Go

	fmt.Println(strings.TrimSpace("  Hello, world!		")) // Hello, world!

	fmt.Println(strings.TrimSuffix("Hello, world!", "orld!")) // Hello, w
	fmt.Println(strings.TrimSuffix("Hello, world!", "wor"))   // Hello, world!
}
