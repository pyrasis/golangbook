package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	s := "Hello, world!"

	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i) // 반복문의 변수를 
                                          // 클로저에서 바로 출력
		}()
	}

	fmt.Scanln()
}
