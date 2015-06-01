package main

import "fmt"

func SumAndDiff(a int, b int) (sum int, diff int) { // 리턴값을 각각 sum, diff로 이름을 정함
	sum = a + b  // 리턴값 변수 sum에 합 대입
	diff = a - b // 리턴값 변수 diff에 차 대입
	return
}

func main() {
	sum, diff := SumAndDiff(6, 2)
	fmt.Println(sum, diff) // 8 4
}
