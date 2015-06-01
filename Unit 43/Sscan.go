package main

import "fmt"

func main() {
	var num1 int
	var num2 float32
	var s string

	input1 := "1\n1.1\nHello"
	n, _ := fmt.Sscan(input1, &num1, &num2, &s) // 공백, 개행 문자로 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n) // 입력 개수: 3
	fmt.Println(num1, num2, s)   // 1 1.1 Hello

	input2 := "1 1.1 Hello"
	n, _ = fmt.Sscanln(input2, &num1, &num2, &s) // 공백으로 구분된 문자열에서 입력을 받음
	fmt.Println("입력 개수:", n)  // 입력 개수: 3
	fmt.Println(num1, num2, s)    // 1 1.1 Hello

	input3 := "1,1.1,Hello"
	n, _ = fmt.Sscanf(input3, "%d,%f,%s", &num1, &num2, &s) // 문자열에서 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n) // 입력 개수: 3
	fmt.Println(num1, num2, s)   // 1 1.1 Hello
}
