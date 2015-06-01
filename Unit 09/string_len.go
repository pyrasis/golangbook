package main

import "fmt"

func main() {
	var s1 string = "한글"
	var s2 string = "Hello"

	fmt.Println(len(s1)) // 6: UTF-8 인코딩의 바이트 길이이므로 6
	fmt.Println(len(s2)) // 5: 알파벳 5글자이므로 5
}
