package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1 string
	s1 = strconv.FormatBool(true)
	fmt.Println(s1) //true: true를 문자열로 변환하여 "true"

	var s2 string
	s2 = strconv.FormatFloat(1.3, 'f', -1, 32) // f, fmt, prec, bitSize
	fmt.Println(s2)                            // 1.3: 1.3을 문자열로 변환하여 "1.3"

	var s3 string
	s3 = strconv.FormatInt(-10, 10)
	fmt.Println(s3) // -10: -10을 10진수로된 문자열로 변환하여 "-10"

	var s4 string
	s4 = strconv.FormatUint(32, 16)
	fmt.Println(s4) // 20: 32를 16진수로된 문자열로 변환하여 "20"
}
