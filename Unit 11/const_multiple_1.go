package main

import "fmt"

func main() {
	const x, y int = 30, 50       // x = 30, y = 50
	const age, name = 10, "Maria" // age = 10, name = "Maria"

	fmt.Println(x, y)
	fmt.Println(age, name)
}
