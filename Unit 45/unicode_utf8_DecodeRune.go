package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("안녕하세요")

	r, size := utf8.DecodeRune(b)
	fmt.Printf("%c %d\n", r, size) // 안 3: "안녕하세요"의 첫 글자를 디코딩하여 '안', 바이트 수 3

	r, size = utf8.DecodeRune(b[3:]) // '안'의 길이가 3이므로 인덱스 3부터 부분 슬라이스를 만들면 "녕하세요"가 됨
	fmt.Printf("%c %d\n", r, size)   // 녕 3: "녕하세요"를 첫 글자를 디코딩하여 '녕', 바이트 수 3

	r, size = utf8.DecodeLastRune(b)
	fmt.Printf("%c %d\n", r, size) // 요 3: "안녕하세요"의 마지막 글자를 디코딩하여 '요', 바이트 수 3

	// '요'의 길이가 3이므로 // 문자열 길이-3을 하여 부분 슬라이스를 만들면
	// "안녕하세"가 됨 
	r, size = utf8.DecodeLastRune(b[:len(b)-3]) 
	                                            
	fmt.Printf("%c %d\n", r, size) // 세 3: "안녕하세"의 마지막 글자를 디코딩하여 '세', 바이트 수 3
}
