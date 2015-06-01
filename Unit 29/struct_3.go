package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func main() {
	var rect1 Rectangle // 구조체 인스턴스 생성
	var rect2 *Rectangle = new(Rectangle) // 구조체 포인터 선언 후 메모리 할당

	rect1.height = 20 // 구조체 인스턴의 필드에 접근할 때 .을 사용
	rect2.height = 62 // 구초체 포인터에 접근할 때도 .을 사용

	fmt.Println(rect1) // {0 20}: 구조체 인스턴스의 값
	fmt.Println(rect2) // &{0 62}: 구조체 포인터이므로 앞에 &가 붙음
}
