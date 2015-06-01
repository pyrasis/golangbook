package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"` // 구조체에 태그 설정
	age  int    `tag1:"나이" tag2:"Age"`  // 구조체에 태그 설정
}

func main() {
	p := Person{}

	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2")) // true 이름 Name

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2")) // true 나이 Age
}
