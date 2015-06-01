package main

/*
#include <stdio.h>

extern int sum(int a, int b); // Go 언어의 함수는 extern으로 선언

static inline void CExample() {
	int r = sum(1, 2); // Go 언어의 sum 함수 호출
	printf("%d\n", r);
}
*/
import "C"

//export sum
func sum(a, b C.int) C.int { // C 언어에서 사용할 수 있도록 매개변수와 리턴값 자료형을 
                             // C 언어용으로 맞춰줌
	return a + b
}

func main() {
	C.CExample()
}
