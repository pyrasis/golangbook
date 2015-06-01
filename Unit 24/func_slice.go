package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	f := []func(int, int) int{sum, diff} // 함수를 저장할 수 있는 슬라이스를 생성한 뒤 함수로 초기화

	fmt.Println(f[0](1, 2)) // 3: 배열의 첫 번째 요소로 함수 호출
	fmt.Println(f[1](1, 2)) // -1: 배열의 두 번째 요소로 함수 호출
}
