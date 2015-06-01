package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	var b [3]int

	b = a          // 배열의 요소가 모두 복사됨
	b[0] = 9       //  b[0]에 9를 대입하면 b의 첫 번째 요소만 바뀜

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [9 2 3]
}