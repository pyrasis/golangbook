package main

/*
extern void goCallback(void *p, int n); // Go 언어의 함수는 extern으로 선언

static inline void CExample(void *p) { // 함수 포인터를 매개변수로 받음
	goCallback(p, 100);            // 받은 함수 포인터를 그대로 넣어줌
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export goCallback
func goCallback(p unsafe.Pointer, n C.int) { // C 언어에서 넘겨준 void *는 
                                             // unsafe.Pointer로 받음
	f := *(*func(C.int))(p) // p를 함수 포인터 타입으로 변환한 뒤 역참조
	f(n)                    // 함수 f 호출
}

func main() {
	f := func(n C.int) { // 매개변수로 C.int를 받는 함수
		fmt.Println(n)
	}

	C.CExample(unsafe.Pointer(&f)) // 함수 f의 포인터를 unsafe.Pointer로 변환하여 넣어줌
}
