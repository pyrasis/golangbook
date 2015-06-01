package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (_ Rectangle) dummy() { // _로 리시버 변수 생략
	fmt.Println("dummy")
}

func main() {
	var r Rectangle
	r.dummy() // dummy
}
