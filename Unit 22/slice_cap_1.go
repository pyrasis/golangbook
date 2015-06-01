package main

import "fmt"

func main() {
	a := make([]int, 5, 10) // 길이가 5이고 용량이 10인 슬라이스 생성

	fmt.Println(len(a)) // 길이는 5
	fmt.Println(cap(a)) // 용량은 10
}