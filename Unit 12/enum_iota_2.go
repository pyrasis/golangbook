package main

import "fmt"

func main() {
	const (
		a = 1 << iota // a == 1 (1 << 0)
		b = 1 << iota // b == 2 (1 << 1)
		c = 1 << iota // c == 4 (1 << 2)
		d = 1 << iota // d == 8 (1 << 3)
	)

	fmt.Println(a, b, c, d)
}
