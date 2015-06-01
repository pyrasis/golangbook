package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	a = append(a, b...) // 슬라이스 a에 슬라이스 b를 붙일 때는 b...을 씀

	fmt.Println(a) // [1 2 3 4 5 6]
}