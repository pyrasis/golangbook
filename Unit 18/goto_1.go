package main

import "fmt"

func main() {
	var a int = 1

	if a == 1 {
		goto ERROR1 // 에러 1 상황이면 ERROR1 레이블로 이동
	}

	if a == 2 {
		goto ERROR2 // 에러 2 상황이면 ERROR2 레이블로 이동
	}

	if a == 3 {
		goto ERROR1 // 에러 1 상황이면 ERROR1 레이블로 이동
	}

	fmt.Println(a)
	fmt.Println("Success") // 정상 실행

	return

ERROR1: // 에러 처리 1
	fmt.Println("Error 1")
	return

ERROR2: // 에러 처리 2
	fmt.Println("Error 2")
	return
}
