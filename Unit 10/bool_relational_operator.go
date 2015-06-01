package main

import "fmt"

func main() {
	var num1 int = 3
	var num2 int = 10

	fmt.Println(num1 > num2)  // false
	fmt.Println(num1 < num2)  // true
	fmt.Println(num1 != num2) // true
	fmt.Println(num1 >= num2) // false
	fmt.Println(num1 <= num2) // true
}
