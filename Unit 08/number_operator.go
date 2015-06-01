package main

import "fmt"

func main() {
	var num1 uint8 = 3
	var num2 uint8 = 2

	fmt.Println(num1 + num2)  // 5
	fmt.Println(num1 - num2)  // 1
	fmt.Println(num1 * num2)  // 6
	fmt.Println(num1 / num2)  // 1
	fmt.Println(num1 % num2)  // 1
	fmt.Println(num1 << num2) // 12
	fmt.Println(num1 >> num2) // 0
	fmt.Println(^num1)        // 252: 비트 반전 연산자
}
