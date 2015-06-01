package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

	fmt.Println(runtime.GOMAXPROCS(0)) // 설정 값 출력

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func(n int) { // 익명 함수를 고루틴으로 실행
			fmt.Println(s, n)
		}(i)
	}

	fmt.Scanln()
}
