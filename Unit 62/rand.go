package main

/*
#include <stdlib.h>
*/
import "C"
import "fmt"

func main() {
	r := C.rand() // stdlib의 rand 함수 호출

	fmt.Println(r) // 결괏값 출력
}
