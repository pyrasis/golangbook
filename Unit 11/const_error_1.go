package main

import "fmt"

func main() {
	const age int = 10
	const name string = "Maria"
	const score int // 컴파일 에러

	age = 20       // 컴파일 에러
	name = "Grace" // 컴파일 에러
}
