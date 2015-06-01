package main

import (
	"fmt"
	"os"
)

func ReadHello() {
	file, err := os.Open("hello.txt")
	defer file.Close() // 지연 호출한 file.Close()가 맨 마지막에 호출됨

	if err != nil {
		fmt.Println(err)
		return // file.Close() 호출
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return // file.Close() 호출
	}

	fmt.Println(string(buf))

	// file.Close() 호출
}

func main() {
	ReadHello()
}
