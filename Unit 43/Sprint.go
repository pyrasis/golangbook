package main

import "fmt"

func main() {
	var s1 string
	s1 = fmt.Sprint(1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만듦
	fmt.Print(s1)

	var s2 string
	s2 = fmt.Sprintln(1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자(\n)를 붙임
	fmt.Print(s2)

	var s3 string
	s3 = fmt.Sprintf("%d %f %s\n", 1, 1.1, "Hello, world!") // 형식을 지정하여 문자열로 만듦
	fmt.Print(s3)
}
