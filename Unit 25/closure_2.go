package main

import "fmt"

func main() {
	a, b := 3, 5

	f := func(x int) int {
		return a*x + b // 함수 바깥의 변수 a, b 사용
	}

	y := f(5)

	fmt.Println(y) // 20
}
