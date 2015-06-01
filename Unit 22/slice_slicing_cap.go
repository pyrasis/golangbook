package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	b := a[0:6:8] // 인덱스 0부터 6까지 가져와서 부분 슬라이스로 만들고 용량을 8로 설정

	fmt.Println(len(b), cap(b)) // 6 8: 길이가 6이며 용량이 8인 슬라이스
	fmt.Println(b)              // [1 2 3 4 5 6]
}
