package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Println(len(a), cap(a)) // 5 5: 길이가 5이며 용량이 5인 슬라이스

	a = append(a, 6, 7)         // 슬라이스 a에 값 6, 7을 추가

	fmt.Println(len(a), cap(a)) // 7 10: 길이가 7이며 용량이 10인 슬라이스, 용량이 늘어남!
}
