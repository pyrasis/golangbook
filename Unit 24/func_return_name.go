package main

import "fmt"

//                      ↓ 리턴 값 변수의 이름을 r로 지정
func sum(a int, b int) (r int) {
	r = a + b // 리턴값 변수 r에 값 대입
	return    // 리턴값 변수를 사용할 때는 return 뒤에 변수를 지정하지 않음
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)
}
