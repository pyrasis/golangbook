package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hello(n int) {
	r := rand.Intn(100)          // 랜덤한 숫자 생성
	time.Sleep(time.Duration(r)) // 랜덤한 시간 동안 대기
	fmt.Println(n)               // n 출력
}

func main() {
	for i := 0; i < 100; i++ { // 100번 반복하여
		go hello(i)        // 고루틴 100개 생성
	}

	fmt.Scanln()
}
