package main

import "fmt"

type hello interface { // 인터페이스 정의
}

func main() {
	var h hello    // 인터페이스 선언
	fmt.Println(h) // <nil>: 빈 인터페이스이므로 nil이 출력됨
}
