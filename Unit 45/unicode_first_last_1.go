package main

import "fmt"

func main() {
	s := "Hello, world!"

	fmt.Printf("%c\n", s[0])        // H: 인덱스 0이 첫 번째 글자
	fmt.Printf("%c\n", s[len(s)-1]) // !: 문자열 길이에서 1을 뺀 인덱스가 마지막 글자
}
