package main

import "fmt"

func sum(a int, b int, c chan int) {
	c <- a + b          // 채널에 a와 b의 합을 보냄
}

func main() {
	c := make(chan int) // int형 채널 생성

	go sum(1, 2, c)     // sum을 고루틴으로 실행한 뒤 채널을 매개변수로 넘겨줌

	n := <-c            // 채널에서 값을 꺼낸 뒤 n에 대입
	fmt.Println(n)      // 3
}
