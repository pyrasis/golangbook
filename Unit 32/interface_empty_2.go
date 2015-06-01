package main

import (
	"fmt"
	"strconv"
)

type Person struct { // Person 구조체 정의
	name string
	age  int
}

func formatString(arg interface{}) string {
	switch arg.(type) {
	case Person:               // arg의 타입이 Person이라면
		p := arg.(Person)  // Person 타입으로 값을 가져옴
		return p.name + " " + strconv.Itoa(p.age) // 각 필드를 합쳐서 리턴
	case *Person:              // arg의 타입이 Person 포인터라면
		p := arg.(*Person) // *Person 타입으로 값을 가져옴
		return p.name + " " + strconv.Itoa(p.age) // 각 필드를 합쳐서 리턴
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(Person{"Maria", 20}))
	fmt.Println(formatString(&Person{"John", 12}))

	var andrew = new(Person)
	andrew.name = "Andrew"
	andrew.age = 35

	fmt.Println(formatString(andrew))
}
