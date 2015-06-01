package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func main() {
	f := map[string]func(int, int) int{ // 함수를 저장할 수 있는 맵을 생성한 뒤 함수로 초기화
		"sum":  sum,
		"diff": diff,
	}

	fmt.Println(f["sum"](1, 2))  // 3: 맵에 sum 키를 지정하여 함수 호출
	fmt.Println(f["diff"](1, 2)) // -1: 맵에 diff 키를 지정하여 함수 호출
}
