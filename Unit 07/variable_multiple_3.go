package main

import "fmt"

func main() {
	var x, y int
	var age int

	x, y, age = 10, 20, 5 // x = 10, y = 20, age = 5: 병렬 할당

	fmt.Println(x, y, age)
}
