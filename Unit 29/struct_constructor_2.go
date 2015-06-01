package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func main() {
	rect := &Rectangle{20, 10} // 구조체를 초기화한 뒤 메모리 주소를 대입

	fmt.Println(rect) // &{20 10}
}
