package main

import (
	"fmt"
	"reflect"
)

type Data struct { // 구조체 정의
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num)) // int: reflect.TypeOf로 자료형 이름 출력

	var s string = "Hello, world!"
	fmt.Println(reflect.TypeOf(s)) // string: reflect.TypeOf로 자료형 이름 출력

	var f float32 = 1.3
	fmt.Println(reflect.TypeOf(f)) // float32: reflect.TypeOf로 자료형 이름 출력

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data)) // main.Data: reflect.TypeOf로 구조체 이름 출력
}
