package main

import "fmt"

//              ↓ 리턴 값이 익명 함수
func calc() func(x int) int {
	a, b := 3, 5           // 지역 변수는 함수가 끝나면 소멸되지만
	return func(x int) int {
		return a*x + b // 클로저이므로 함수를 호출 할 때마다 변수 a와 b의 값을 사용할 수 있음
	}
	// ↑ 익명 함수를 리턴
}

func main() {
	f := calc() // calc 함수를 실행하여 리턴값으로 나온 클로저를 변수에 저장

	fmt.Println(f(1)) // 8
	fmt.Println(f(2)) // 11
	fmt.Println(f(3)) // 14
	fmt.Println(f(4)) // 17
	fmt.Println(f(5)) // 20
}
