package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data = []int{} // int형 슬라이스 생성

	go func() {                             // 고루틴에서
		for i := 0; i < 1000; i++ {     // 1000번 반복하면서
			data = append(data, 1)  // data 슬라이스에 1을 추가

			runtime.Gosched()       // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() {                             // 고루틴에서
		for i := 0; i < 1000; i++ {     // 1000번 반복하면서
			data = append(data, 1)  // data 슬라이스에 1을 추가

			runtime.Gosched()       // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second)      // 2초 대기

	fmt.Println(len(data))           // data 슬라이스의 길이 출력
}
