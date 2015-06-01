package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "안녕하세요"

	r, _ := utf8.DecodeRuneInString(s) // UTF-8 문자열의 첫 글자와 바이트 수를 리턴
	fmt.Printf("%c\n", r)              // 안: 문자열의 첫 글자

	r, _ = utf8.DecodeLastRuneInString(s) // UTF-8 문자열의 첫 글자와 바이트 수를 리턴
	fmt.Printf("%c\n", r)                 // 요: 문자열의 마지막 글자
}
