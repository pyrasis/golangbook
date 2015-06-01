package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // CPU를 하나만 사용

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) {          // 익명 함수를 고루틴으로 실행(클로저)
			fmt.Println(s, n) // s와 매개변수로 받은 n 값 출력
		}(i)                      // 반복문의 변수는 매개변수로 넘겨줌
	}

	fmt.Scanln()
}
