package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))               // *int
	fmt.Println(reflect.ValueOf(a))              // <*int Value>
	fmt.Println(reflect.ValueOf(a).Int())        // 런타임 에러
	fmt.Println(reflect.ValueOf(a).Elem())       // <int Value>
	fmt.Println(reflect.ValueOf(a).Elem().Int()) // 1

	var b interface{}
	b = 1

	fmt.Println(reflect.TypeOf(b))         // int
	fmt.Println(reflect.ValueOf(b))        // <int Value>
	fmt.Println(reflect.ValueOf(b).Int())  // 1
	fmt.Println(reflect.ValueOf(b).Elem()) // 런타임 에러
}
