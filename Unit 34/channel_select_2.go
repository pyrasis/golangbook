package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)    // int형 채널 생성
	c2 := make(chan string) // string 채널 생성

	go func() {
		for {
			i := <-c1                          // 채널 c1에서 값을 꺼낸 뒤 i에 대입
			fmt.Println("c1 :", i)             // i 값을 출력
			time.Sleep(100 * time.Millisecond) // 100 밀리초 대기
		}
	}()

	go func() {
		for {
			c2 <- "Hello, world!"              // 채널 c2에 Hello, world!를 보냄
			time.Sleep(500 * time.Millisecond) // 500 밀리초 대기
		}
	}()

	go func() {
		for { // 무한 루프
			select {
			case c1 <- 10:                 // 매번 채널 c1에 10을 보냄
			case s := <-c2:                // c2에 값이 들어왔을 때는 값을 꺼낸 뒤 s에 대입
				fmt.Println("c2 :", s) // s 값을 출력
			}
		}
	}()

	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}
