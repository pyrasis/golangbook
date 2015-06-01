package main

import (
	"fmt"
	"sync"
	"time"
)

// 실제 작업을 처리하는 worker 함수
func worker(n int, done <-chan struct{}, jobs <-chan int, c chan<- string) {
	for j := range jobs { // 작업 요청 채널에서 값을 가져와서 실행
		select {
		case c <- fmt.Sprintf("Worker: %d, Job: %d", n, j):
		case <-done: // 채널 done이 닫히면 worker 함수를 빠져나옴
			return
		}
	}
}

func main() {
	jobs := make(chan int)      // 작업을 요청할 채널
	done := make(chan struct{}) // 작업 고루틴에 정지 신호를 보낼 채널
	c := make(chan string)      // 결괏값을 저장할 채널

	var wg sync.WaitGroup
	const numWorkers = 5
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ { // 작업을 처리할 고루틴을 5개 생성
		go func(n int) {
			worker(n, done, jobs, c)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait() // 고루틴이 끝날 때까지 대기
		close(c)  // 대기가 끝나면 결괏값 채널을 닫음
	}()

	go func() {
		for i := 0; i < 10; i++ { // 작업 요청을 10개 생성하여
			jobs <- i         // jobs 채널에 보냄
			time.Sleep(10 * time.Millisecond)
		}
		close(done) // 채널을 닫아서 모든 worker 고루틴 정지
		close(jobs) // worker 함수의 range 대기도 종료
	}()

	for r := range c {     // 결과 채널에 값이 들어올 때까지 대기한 뒤 값을 가져옴
		fmt.Println(r) // 결괏값 출력
	}
}
