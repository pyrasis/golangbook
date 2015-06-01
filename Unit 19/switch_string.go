package main

import "fmt"

func main() {
	s := "world"

	switch s {                   // 값을 판단할 변수 설정
	case "hello":                // 각 조건에 일치하는
		fmt.Println("hello") // 코드를 실행합니다.
	case "world":                // 문자열 "world"와 변수의 값이 일치하므로
		fmt.Println("world") // 이 부분을 실행하고 이후 실행을 중단
	default:
		fmt.Println("일치하는 문자열이 없습니다.")
	}
}