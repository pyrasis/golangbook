package main

import "fmt"

type Person struct { // 사람 구조체 정의
	name string
	age  int
}

func (p *Person) greeting() { // 인사(greeting) 함수 작성
	fmt.Println("Hello~")
}

type Student struct {
	p      Person // 학생 구조체 안에 사람 구조체를 필드로 가지고 있음. Has-a 관계
	school string
	grade  int
}

func main() {
	var s Student

	s.p.greeting() // Hello~
}
