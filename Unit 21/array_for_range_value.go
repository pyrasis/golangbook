package main

import "fmt"

func main() {
	a := [5]int{32, 29, 78, 16, 81}

	for _, value := range a { // 인덱스는 생략, value에 배열 요소의 값이 들어감
		fmt.Println(value)
	}
}