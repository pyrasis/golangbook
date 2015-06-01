package main

import "fmt"

func hello(n *int) {
	*n = 2 // 포인터 변수 n를 역참조하여 메모리에 2를 대입
}

func main() {
	var n int = 1

	hello(&n) // 1이 들어있는 변수 n의 메모리 주소를 hello 함수에 넘김

	fmt.Println(n) // 2: hello 함수에서 n의 메모리 공간에 2를 대입했으므로 바깥에 있는 n의 값이 바뀌었음
}
