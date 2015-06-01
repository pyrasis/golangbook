package main

import "fmt"

func main() {
	var num1, num2 int
	n, _ := fmt.Scanf("%d,%d", &num1, &num2) // 정수형으로 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2)
}
