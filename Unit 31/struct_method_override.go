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
	Person
	school string
	grade  int
}

func (p *Student) greeting() { // 이름이 같은 메서드를 정의하면 오버라이딩됨
	fmt.Println("Hello Students~")
}

func main() {
	var s Student

	s.Person.greeting() // Hello~
	s.greeting()        // Hello Students~
}
