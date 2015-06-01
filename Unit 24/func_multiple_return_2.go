package main

import "fmt"

func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b // 리턴할 값 또는 변수를 차례대로 나열
}

func main() {
	_, diff := SumAndDiff(6, 2) // 첫 번째 리턴값은 _으로 생략, 두 번째 리턴값만 사용
	fmt.Println(diff)           // 4: 차
}
