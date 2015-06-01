package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := []string{"Hello,", "world!"}
	fmt.Println(strings.Join(s1, " ")) // Hello, world!

	s2 := strings.Split("Hello, world!", " ")
	fmt.Println(s2[1]) // world!

	s3 := strings.Fields("Hello, world!")
	fmt.Println(s3[1]) // world!

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r) // r이 한글 유니코드이면 true를 리턴
	}
	s4 := strings.FieldsFunc("Hello안녕Hello", f)
	fmt.Println(s4) // [Hello Hello]

	fmt.Println(strings.Repeat("Hello", 2)) // HelloHello

	fmt.Println(strings.Replace("Hello, world!", "world", "Go", 1)) // Hello, Go!
	fmt.Println(strings.Replace("Hello Hello", "llo", "Go", 2))     // HeGo HeGo
}
