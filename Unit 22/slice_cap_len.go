package main

import "fmt"

func main() {
	a := make([]int, 5, 10)

	fmt.Println(len(a)) // 길이는 5
	fmt.Println(cap(a)) // 용량은 10
}