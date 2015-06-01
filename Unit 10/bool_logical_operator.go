package main

import "fmt"

func main() {
	fmt.Println(true && true)   // true
	fmt.Println(true && false)  // false
	fmt.Println(false && false) // false
	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(false || false) // false
	fmt.Println(!true)          // false
	fmt.Println(!false)         // true
}
