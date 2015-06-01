package main

import "fmt"

func main() {
	a := [5]int{32, 29, 78, 16, 81}

	for value := range a { // value에는 값 대신 인덱스가 들어감
		fmt.Println(value)
	}
}