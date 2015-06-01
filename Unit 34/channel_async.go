package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2) // 버퍼가 2개인 비동기 채널 생성
	count := 4                 // 반복할 횟수

	go func() {
		for i := 0; i < count; i++ {
			done <- true                // 채널에 true를 보냄, 버퍼가 가득차면 대기
			fmt.Println("고루틴 : ", i) // 반복문의 변수 출력
		}
	}()

	for i := 0; i < count; i++ {
		<-done                         // 버퍼에 값이 없으면 대기, 값을 꺼냄
		fmt.Println("메인 함수 : ", i) // 반복문의 변수 출력
	}
}
