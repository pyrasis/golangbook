package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	wg := new(sync.WaitGroup) // 대기 그룹 생성

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done() // 고루틴이 끝나기 직전에 wg.Done 함수 호출
			fmt.Println(n)
		}(i)
	}

	wg.Wait()
	fmt.Println("the end")
}
