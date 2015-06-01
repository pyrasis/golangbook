package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() {
	fmt.Println("Hello~")
}

type Student struct {
	Person // 필드명 없이 타입만 선언하면 포함(Is-a) 관계가 됨
	school string
	grade  int
}

func main() {
	var s Student

	s.Person.greeting() // Hello~
	s.greeting()        // Hello~
}
