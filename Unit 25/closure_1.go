package main

import "fmt"

//    ↓ 함수 안에서
func main() {
	//      ↓ 익명 함수를 선언 및 정의
	sum := func(a, b int) int {
		return a + b
	}

	r := sum(1, 2) // 익명함수 사용

	fmt.Println(r) // 3
}
