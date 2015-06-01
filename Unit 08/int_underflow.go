package main

import "fmt"

func main() {
	var num1 uint8 = 0 - 1  // 오버플로우 컴파일 에러 발생
	var num2 uint16 = 0 - 1 // 오버플로우 컴파일 에러 발생
	var num3 uint32 = 0 - 1 // 오버플로우 컴파일 에러 발생
	var num4 uint64 = 0 - 1 // 오버플로우 컴파일 에러 발생
}
