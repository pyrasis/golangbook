package main

import "fmt"

func main() {
	c := make(chan int, 1)

	go func() {
		c <- 1
	}()

	a, ok := <-c       // 채널이 닫혔는지 확인
	fmt.Println(a, ok) // 1 true: 채널은 열려 있고 1을 꺼냄

	close(c)           // 채널을 닫음
	a, ok = <-c        // 채널이 닫혔는지 확인
	fmt.Println(a, ok) // 0 false: 채널이 닫혔음
}
