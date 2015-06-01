package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {                       // 고루틴 2,000개 생성
			atomic.AddInt32(&data, 1) // 원자적 연산으로 1씩 더함
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {                        // 고루틴 1,000개 생성
			atomic.AddInt32(&data, -1) // 원자적 연산으로 1씩 뺌
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)
}
