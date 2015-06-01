package main

import "fmt"

func main() {
	var numPtr *int = new(int)

	numPtr++              // 컴파일 에러. 포인터 연산은 허용하지 않음
	numPtr = 0xc0820062d0 // 컴파일 에러. 메모리 주소를 직접 대입할 수 없음

	fmt.Println(numPtr)
}
