package main

import "fmt"

func main() {
	a := []int{1, 2, 3}

	a = append(a, 4, 5, 6)

	fmt.Println(a) // [1 2 3 4 5 6]
}