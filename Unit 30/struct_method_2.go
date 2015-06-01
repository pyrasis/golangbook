package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

//           ↓ 포인터 방식
func (rect *Rectangle) scaleA(factor int) {
	rect.width = rect.width * factor   // 포인터이므로 원래의 값이 변경됨
	rect.height = rect.height * factor // 포인터이므로 원래의 값이 변경됨
}

//          ↓ 일반 구조체 방식
func (rect Rectangle) scaleB(factor int) {
	rect.width = rect.width * factor   // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
	rect.height = rect.height * factor // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
}

func main() {
	rect1 := Rectangle{30, 30}
	rect1.scaleA(10)
	fmt.Println(rect1) // {300 300}: rect1에 바뀐 값이 들어감

	rect2 := Rectangle{30, 30}
	rect2.scaleB(10)
	fmt.Println(rect2) // {30 30}: rect2는 값이 바뀌지 않음
}
