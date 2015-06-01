package main

import "fmt"

func main() {
	var (
		x, y      int = 30, 50      // x와 y는 int형으로 결정
		age, name     = 10, "Maria" // age는 int, name은 string으로 결정
	)

	fmt.Println(x, y, age, name)
}
