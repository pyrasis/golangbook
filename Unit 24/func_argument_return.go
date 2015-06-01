package main

import "fmt"

func sum(a int, b int) int { // int형 매개변수 a, b 그리고 int 형 리턴값을 가지는 함수 정의
	return a + b
}

func main() {
	r := sum(1, 2)
	fmt.Println(r) // 3
}
