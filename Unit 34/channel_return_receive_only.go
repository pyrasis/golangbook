package main

import "fmt"

//                    ↓ 함수의 리턴 값은 int 형 받기 전용 채널
func sum(a, b int) <-chan int {
	out := make(chan int) // 채널 생성
	go func() {
		out <- a + b // 채널에 a와 b의 합을 보냄
	}()
	return out           // 채널 변수 자체를 리턴
}

func main() {
	c := sum(1, 2)   // 채널을 리턴값으로 받아서 c에 대입

	fmt.Println(<-c) // 3: 채널에서 값을 꺼냄
}
