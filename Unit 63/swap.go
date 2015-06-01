package main

/*
extern void CExample();                       // Go 언어에서 C 언어의 함수를 호출하기 위한 선언
extern struct swap_return swap(int a, int b); // C언어에서 Go 언어의 함수를 호출하기 위한 선언.
                                              // 구조체를 이용하여 두 개의 값을 리턴
*/
import "C"

//export swap
func swap(a, b C.int) (C.int, C.int) { // C 언어에서 사용할 함수, 리턴값이 두 개
	return b, a
}

func main() {
	C.CExample()
}
