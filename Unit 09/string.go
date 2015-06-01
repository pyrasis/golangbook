package main

import "fmt"

func main() {
	var s1 string = "Hello, world!\n"
	var s2 string = "안녕하세요\n"
	var s3 string = "\ud55c\uae00"             // 한글: 유니코드 문자 코드로 저장
	var s4 string = "\U0000d55c\U0000ae00"     // 한글: 유니코드 문자 코드로 저장
	var s5 string = "\xed\x95\x9c\xea\xb8\x80" // 한글: UTF-8 인코딩의 바이트 값으로 저장

	var s7 string = `안녕하세요
	Hello, world!`

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s7)
}
