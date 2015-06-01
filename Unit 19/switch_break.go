package main

import "fmt"

func main() {
	s := "Hello"
	i := 2

	switch i {                             // 값을 판단할 변수 설정
	case 1:
		fmt.Println(1)
	case 2:                                // i가 2이고
		if s == "Hello" {              // s가 "Hello"이면
			fmt.Println("Hello 2") // Hello 2를 출력하고
			break                  // switch 분기문 실행을 중단
		}

		fmt.Println(2)
	}
}