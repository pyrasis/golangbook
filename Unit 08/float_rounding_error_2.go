package main

import "fmt"
import "math"

func main() {
	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a) // 9.000000000000004

	const epsilon = 1e-14                   // Go 언어 머신 입실론
	fmt.Println(math.Abs(a-9.0) <= epsilon) // true: 연산한 값과 비교할 값의 차이를 구한 뒤
	                                        // 머신 입실론보다 작거나 같은지 비교
}
