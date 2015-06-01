package main

import "fmt"

func main() {
	func() { // 함수에 이름이 없음
		fmt.Println("Hello, world!")
	}()

	func(s string) {   // 익명 함수를 정의한 뒤
		fmt.Println(s)
	}("Hello, world!") // 바로 호출

	r := func(a int, b int) int { // 익명 함수를 정의한 뒤
		return a + b
	}(1, 2)                       // 바로 호출하여 리턴값을 변수 r에 저장

	fmt.Println(r) // 3
}
