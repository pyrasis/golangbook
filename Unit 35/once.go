package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello() {
	fmt.Println("Hello, world!")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	once := new(sync.Once) // Once 생성

	for i := 0; i < 3; i++ {
		go func(n int) {   // 고루틴 3개 생성
			fmt.Println("goroutine : ", n)

			once.Do(hello) // 고루틴은 3개지만 hello 함수를 한 번만 실행
		}(i)
	}

	fmt.Scanln()
}
