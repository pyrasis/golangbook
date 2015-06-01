package main

import (
	"fmt"
	"log"
)

func helloOne(n int) (string, error) {
	if n == 1 {                           // 1일때만
		s := fmt.Sprintln("Hello", n) // Hello 문자열을 리턴
		return s, nil                 // 정상 동작이므로 에러 값은 nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n) // 1이 아닐 때는 에러를 리턴
}

func main() {
	s, err := helloOne(1) // 매개변수에 1을 넣었으므로 정상 동작
	if err != nil {
		log.Print(err) // 에러가 없으므로 실행되지 않음
	}

	fmt.Println(s) // Hello 1

	s, err = helloOne(2)   // 매개변수에 2를 넣었으므로 에러 발생
	if err != nil {
		log.Print(err) // 에러 문자열이 출력됨
	}

	// long.Print 함수는 프로그램을 종료하지 않으므로 이 아래는 계속 실행됨
	fmt.Println(s)

	fmt.Println("Hello, world!")
}
