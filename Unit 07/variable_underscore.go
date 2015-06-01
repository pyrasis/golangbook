package main

import "fmt"

func main() {
	a := 1
	b := 2

	_ = b // 사용하지 않는 변수로 인한 컴파일 에러 방지

	fmt.Println(a)
}
