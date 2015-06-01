package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := make([]int, 3)

	copy(b, a)     // 슬라이스를 복사하였으므로
	b[0] = 9       // b[0]만 바뀌고, a[0]은 바뀌지 않음

	fmt.Println(a) // [1 2 3]
	fmt.Println(b) // [9 2 3]
}
