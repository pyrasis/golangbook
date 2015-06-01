package main

import "fmt"

func main() {
	var numPtr *int = new(int) // new 함수로 공간 할당

	*numPtr = 1          // 역참조로 포인터형 변수에 값을 대입

	fmt.Println(*numPtr) // 1: 포인터형 변수에서 값을 가져오기
}
