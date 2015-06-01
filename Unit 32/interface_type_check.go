package main

import (
	"fmt"
	"strconv"
)

type Person struct { // Person 구조체 정의
	name string
	age  int
}

func main() {
	var t interface{}
	t = Person{"Maria", 20}

	if v, ok := t.(Person); ok {
		fmt.Println(v, ok)
	}
}
