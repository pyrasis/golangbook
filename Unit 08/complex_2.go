package main

import "fmt"

func main() {
	var num1 complex64 = complex(1, 2)                    // 실수부 1, 허수부 2i
	var num2 complex128 = complex(5.32521e-10, .12345E+2) // 실수부 부동소수점 5.32521e-10,
	                                                      // 허수부 부동소수점 .12345E+2i

	fmt.Println(num1)
	fmt.Println(num2)
}
