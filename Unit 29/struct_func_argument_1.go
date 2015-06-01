package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func rectangleArea(rect *Rectangle) int { // 매개변수로 구조체 포인터를 받음
	return rect.width * rect.height
}

func main() {
	rect := Rectangle{30, 30}

	area := rectangleArea(&rect) // 구조체의 포인터를 넘김

	fmt.Println(area) // 900
}
