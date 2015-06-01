package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, world!"
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따르는 
                                  // 읽기 인스턴스 r 생성

	var s1, s2 string
	n, _ := fmt.Fscanf(r, "%s %s", &s1, &s2) // 형식을 지정하여 읽기 인스턴스 r에서 
                                                 // 문자열 읽기

	fmt.Println("입력 개수:", n)             // 입력 개수: 2
	fmt.Println(s1)                          // Hello,
	fmt.Println(s2)                          // world!
}
