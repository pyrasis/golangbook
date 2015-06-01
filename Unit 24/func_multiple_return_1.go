package main

import "fmt"

func SumAndDiff(a int, b int) (int, int) { // int형 리턴값이 두 개인 함수 정의
	return a + b, a - b // 리턴할 값 또는 변수를 차례대로 나열. 합과 차를 동시에 리턴
}

func main() {
	sum, diff := SumAndDiff(6, 2) // 변수 두 개에 리턴값 두 개를 대입
	fmt.Println(sum, diff)        // 8 4: 합과 차
}
