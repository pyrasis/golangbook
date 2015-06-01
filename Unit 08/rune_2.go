package main

import "fmt"

func main() {
	var r1 rune = "한"   // 컴파일 에러
	var r2 rune = '한글' // 컴파일 에러
	var r3 rune = "한글" // 컴파일 에러
}
