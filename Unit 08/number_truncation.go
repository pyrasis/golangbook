package main

import "fmt"

func main() {
	var num1 uint16 = 10
	var num2 uint32 = 80000

	fmt.Println(num1 + uint16(num2)) // 14474: uint32에서 uint16으로 변환하면서 넘치는 값을 버림
}
