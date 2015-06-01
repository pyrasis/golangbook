package main

import "fmt"

type MyInt int // int형을 MyInt로 정의

func (i MyInt) Print() { // MyInt에 Print 메서드를 연결
	fmt.Println(i)
}

type Printer interface { // Print 메서드를 가지는 인터페이스 정의
	Print()
}

func main() {
	var i MyInt = 5

	var p Printer // 인터페이스 선언

	p = i     // i를 인터페이스 p에 대입
	p.Print() // 5: 인터페이스를 통하여 MyInt의 Print 메서드 호출
}
