package main

import "fmt"

func main() {
	a := []int{32, 29, 78, 16, 81} // 슬라이스를 생성하면서 값을 초기화

	b := []int{
		32,
		29,
		78,
		16,
		81,  // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}

	fmt.Println(a)
	fmt.Println(b)
}