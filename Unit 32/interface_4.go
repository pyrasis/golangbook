package main

import "fmt"

type MyInt int // int 형을 MyInt로 정의

func (i MyInt) Print() { // MyInt에 Print 메서드를 연결
	fmt.Println(i)
}

type Rectangle struct { // 사각형 구조체 정의
	width, height int
}

func (r Rectangle) Print() { // Rectangle에 Print 메서드를 연결
	fmt.Println(r.width, r.height)
}

type Printer interface { // Print 메서드를 가지는 인터페이스 정의
	Print()
}

func main() {
	var i MyInt = 5
	r := Rectangle{10, 20}

	pArr := []Printer{i, r} // 슬라이스 형태로 인터페이스 초기화
	for index, _ := range pArr {
		pArr[index].Print() // 슬라이스를 순회하면서 Print 메서드 호출
	}

	for _, value := range pArr {
		value.Print() // 슬라이스를 순회하면서 Print 메서드 호출
	}
}
