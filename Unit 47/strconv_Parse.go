package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error

	var b1 bool
	b1, err = strconv.ParseBool("true")
	fmt.Println(b1, err) // true <nil>: 문자열로 "true"를 불로 변환하여 true

	var num1 float64
	num1, err = strconv.ParseFloat("1.3", 64)
	fmt.Println(num1, err) // 1.3 <nil>: 문자열 "1.3"을 실수로 변환하여 1.3

	var num2 int64
	num2, err = strconv.ParseInt("-10", 10, 32)
	fmt.Println(num2, err) // -10 <nil>: 10진수로된 문자열 "-10"을 정수로 변환하여 -10

	var num3 uint64
	num3, err = strconv.ParseUint("20", 16, 32)
	fmt.Println(num3, err) // 32: 16진수로된 문자열 "20"을 정수로 변환하여 32
}
