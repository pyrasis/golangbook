package main

import "fmt"
import "unsafe" // Sizeof 함수를 사용하기 unsafe 패키지 사용

func main() {
	var num1 int8  = 1
	var num2 int16 = 1
	var num3 int32 = 1
	var num4 int64 = 1

	fmt.Println(unsafe.Sizeof(num1)) // 1
	fmt.Println(unsafe.Sizeof(num2)) // 2
	fmt.Println(unsafe.Sizeof(num3)) // 4
	fmt.Println(unsafe.Sizeof(num4)) // 8
}
