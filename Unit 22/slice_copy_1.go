package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 3) // make 함수로 공간을 할당

	copy(b, a)     // 슬라이스 a의 요소를 슬라이스 b에 복사

	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(b) // [1 2 3]: 슬라이스 b의 길이는 3이므로 a의 요소 3개만 복사됨
}
