package main

import "fmt"

func main() {
	var num1 int = 3
	var num2 float32 = 2.2

	fmt.Println(num1 + num2)           // 정수 + 실수이므로 컴파일 에러
	fmt.Println(float32(num1) + num2)  // 5.2: 정수를 실수로 변환
	fmt.Println(num1 + int(num2))      // 5: 실수를 정수로 변환하여 0.2를 버림
}
