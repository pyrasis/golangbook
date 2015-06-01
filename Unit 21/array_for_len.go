package main

import "fmt"

func main() {
	a := [5]int{32, 29, 78, 16, 81}

	for i := 0; i < len(a); i++ { // len 함수로 배열의 길이를 구한 뒤 배열의 길이 만큼 반복
		fmt.Println(a[i])
	}

	fmt.Println(a)
}