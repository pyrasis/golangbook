package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func main() {
	var rect1 *Rectangle   // 구조체 포인터 선언
	rect1 = new(Rectangle) // 구조체 포인터에 메모리 할당

	rect2 := new(Rectangle) // 구조체 포인터 선언과 동시에 메모리 할당

	fmt.Println(rect1)
	fmt.Println(rect2)
}
