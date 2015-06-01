package main

import "fmt"

func main() {
	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)        // 9.000000000000004: 반올림 오차 발생
	fmt.Println(a == 9.0) // false: 실수는 ==로 비교하면 안됨
}
