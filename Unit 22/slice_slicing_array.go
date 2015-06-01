package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5} // 배열 선언

	b := a[:2]     // 배열 a의 일부를 부분 슬라이스로 참조
	b[0] = 9       // 부분 슬라이스는 참조이므로 a[0], b[0]의 값이 모두 바뀜

	fmt.Println(a) // [9 2 3 4 5]
	fmt.Println(b) // [9 2]
}
