package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func rectangleScaleA(rect *Rectangle, factor int) { // 매개변수로 구조체 포인터를 받음
	rect.width = rect.width * factor   // 포인터이므로 원래의 값이 변경됨
	rect.height = rect.height * factor // 포인터이므로 원래의 값이 변경됨
}

func rectangleScaleB(rect Rectangle, factor int) { // 매개변수로 구조체 인스턴스를 받음
	rect.width = rect.width * factor   // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
	rect.height = rect.height * factor // 값이 복사되었으므로 원래의 값에는 영향을 미치지 않음
}

func main() {
	rect1 := Rectangle{30, 30}
	rectangleScaleA(&rect1, 10) // 구조체의 포인터를 넘김
	fmt.Println(rect1)          // {300 300}: rect1에 바뀐 값이 들어감

	rect2 := Rectangle{30, 30}
	rectangleScaleB(rect2, 10) // 구조체 인스턴스를 그대로 넘김
	fmt.Println(rect2)         // {30 30}: rect2는 값이 바뀌지 않음
}
