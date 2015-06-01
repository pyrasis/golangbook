package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var mutex = new(sync.Mutex)    // 뮤텍스 생성
	var cond = sync.NewCond(mutex) // 뮤텍스를 이용하여 조건 변수 생성

	c := make(chan bool, 3) // 비동기 채널 생성

	for i := 0; i < 3; i++ {
		go func(n int) {                        // 고루틴 3개 생성
			mutex.Lock()                    // 뮤텍스 잠금, cond.Wait() 보호 시작
			c <- true                       // 채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait()                     // 조건 변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock()                  // 뮤텍스 잠금 해제, cond.Wait() 보호 종료

		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c // 채널에서 값을 꺼냄, 고루틴 3개가 모두 실행될 때까지 기다림
	}

	mutex.Lock()             // 뮤텍스 잠금, cond.Broadcast() 보호 시작
	fmt.Println("broadcast")
	cond.Broadcast()         // 대기하고 있는 모든 고루틴을 깨움
	mutex.Unlock()           // 뮤텍스 잠금 해제, cond.Signal() 보고 종료

	fmt.Scanln()
}
