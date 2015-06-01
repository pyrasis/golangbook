package main

import "fmt"

func sum(n ...int) int { // int형 가변인자를 받는 함수 정의
	total := 0
	for _, value := range n { // range로 가변인자의 모든 값을 꺼냄
		total += value   // 꺼낸 값을 모두 더함
	}

	return total
}

func main() {
	r := sum(1, 2, 3, 4, 5)
	fmt.Println(r) // 15
}
