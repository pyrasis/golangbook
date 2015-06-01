#include <stdio.h>
#include "_cgo_export.h" // 반드시 포함

void CExample() {
	struct swap_return s;           // 구조체로 리턴값을 받을 변수를 선언
	s = swap(1, 2);                 // Go 언어의 swap 함수 호출
	printf("%d %d\n", s.r0, s.r1);  // 첫 번째 리턴값은 r0, 두 번째 리턴값은 r1
}
