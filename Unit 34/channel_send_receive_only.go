package main

import "fmt"

func producer(c chan<- int) { // 보내기 전용 채널
	for i := 0; i < 5; i++ {
		c <- i
	}

	c <- 100           // 채널에 값을 보냄

	//fmt.Println(<-c) // 채널에서 값을 꺼내면 컴파일 에러
}

func consumer(c <-chan int) { // 받기 전용 채널
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(<-c) // 채널에 값을 꺼냄

	// c <- 1        // 채널에 값을 보내면 컴파일 에러
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
