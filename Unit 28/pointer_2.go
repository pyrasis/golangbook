package main

import "fmt"

func main() {
	var numPtr *int = new(int)

	fmt.Println(numPtr) // 0xc0820062d0: 메모리 주소. 시스템 마다, 실행할 때마다 달라짐
}
