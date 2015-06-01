package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 string = "안녕하세요"
	fmt.Println(utf8.RuneCountInString(s1)) // 5: "안녕하세요"의 실제 길이는 5
}
