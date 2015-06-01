package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error

	var num1 int
	num1, err = strconv.Atoi("100") // 문자열 "100"을 숫자 100으로 변환
	fmt.Println(num1, err)          // 100 <nil>

	var num2 int
	num2, err = strconv.Atoi("10t") // 10t는 숫자가 아니므로 에러가 발생함
	fmt.Println(num2, err)          // 0 strconv.ParseInt: parsing "10t": invalid syntax

	var s string
	s = strconv.Itoa(50) // 숫자 50을 문자열 "50"으로 변환
	fmt.Println(s)       // 50
}
