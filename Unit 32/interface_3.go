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

	var p Printer // 인터페이스 선언

	p = i     // i를 인터페이스 p에 대입
	p.Print() // 5: 인터페이스 p를 통하여 MyInt의 Print 메서드 호출

	p = r     // r을 인터페이스 p에 대입
	p.Print() // 10 20: 인터페이스 p를 통하여 Rectangle의 Print 메서드 호출
}
