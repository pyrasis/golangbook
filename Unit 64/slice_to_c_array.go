package main

/*
#include <stdio.h>

void CExample(void *p) {     // 슬라이스를 void * 타입으로 받음
	char *a = (char *)p; // char * 타입으로 변환

	printf("%c\n", a[0]); // H
	printf("%s\n", a);    // Hello, world!
}
*/
import "C"
import "unsafe"

func main() {
	b := []byte("Hello, world!")      // 바이트 형식으로 슬라이스 선언

	C.CExample(unsafe.Pointer(&b[0])) // 슬라이스 첫 번째 요소의 포인터를 unsafe.Pointer로 변환하여 넣어줌
}
