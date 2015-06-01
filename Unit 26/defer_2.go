package main

import "fmt"

func HelloWorld() {
	//↓ HelloWorld() 함수가 끝나기 직전에 호출
	defer func() {
		fmt.Println("world")
	}()

	func() {
		fmt.Println("Hello")
	}()
}

func main() {
	HelloWorld()
}
