package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value { // 매개변수와 리턴값은 반드시 
                                               // []reflect.Value를 사용
	fmt.Println("Hello, world!")
	return nil
}

func main() {
	var hello func() // 함수를 담을 변수 선언

	fn := reflect.ValueOf(&hello).Elem() // hello의 주소를 넘긴 뒤 Elem으로 값 정보를 가져옴

	v := reflect.MakeFunc(fn.Type(), h) // h의 함수 정보를 생성

	fn.Set(v) // hello의 값 정보인 fn에 h의 함수 정보 v를 설정하여 함수를 연결

	hello()
}
