package main

import "fmt"

func hello() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5 // 리턴할 값을 차례대로 나열
}

func main() {
	a, _, c, _, e := hello() // 2, 4번째 리턴값은 생략
	fmt.Println(a, c, e)     // 1 3 5
}
