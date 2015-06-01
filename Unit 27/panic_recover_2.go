package main

import "fmt"

func f() {
	defer func() {
		s := recover()   // recover 함수로 런타임 에러(패닉) 상황을 복구
		fmt.Println(s)
	}()

	a := [...]int{1, 2}      // 정수가 2개 저장된 배열

	for i := 0; i < 5; i++ { // 배열 크기를 벗어난 접근
		fmt.Println(a[i])
	}
}

func main() {
	f()

	fmt.Println("Hello, world!") // 런타임 에러가 발생했지만 recover 함수로 복구되었기 때문에 이 부분은 정상적으로 실행됨
}
