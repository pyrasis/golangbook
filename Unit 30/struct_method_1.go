package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

//          ↓ 리시버 변수 정의(연결할 구조체 지정)
func (rect *Rectangle) area() int {
	return rect.width * rect.height
	//     ↑  리시버 변수를 사용하여 현재 인스턴스에 접근할 수 있음
}

func main() {
	rect := Rectangle{10, 20}

	fmt.Println(rect.area()) // 200
}
