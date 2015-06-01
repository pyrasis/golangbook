package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	b := a[0:5]    // a의 인덱스 0부터 5까지 참조

	fmt.Println(a) // [1 2 3 4 5]
	fmt.Println(b) // [1 2 3 4 5]
}
