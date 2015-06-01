package main

/*
#include <stdlib.h>

typedef struct _PERSON {
	char *name; // 문자열 포인터
	int age;    // 정수
} PERSON;

PERSON* create(char *name, int age) // 메모리 할당 함수 작성
{
	PERSON *p = (PERSON *)malloc(sizeof(PERSON)); // PERSON 크기로 메모리 할당

	p->name = name; // 값 설정
	p->age = age;   // 값 설정

	return p; // 할당한 메모리 리턴
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var p *C.PERSON

	name := C.CString("Maria")
	age := C.int(20)

	p = C.create(name, age)

	fmt.Println(C.GoString(p.name)) // Maria: char *는 C.GoString 함수로 변환하여 사용
	fmt.Println(p.age)              // 20

	C.free(unsafe.Pointer(name)) // C.CString으로 만든 문자열은 반드시 해제
	C.free(unsafe.Pointer(p))    // C 언어의 malloc 함수로 할당한 메모리는 반드시 해제
}
