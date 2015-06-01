package main

import "fmt"

func main() {
	for i, j := 0, 0; i < 10; i++, j+=2 { // 컴파일 에러. syntax error: unexpected comma, expecting {
		fmt.Println(i, j)
	}
}
