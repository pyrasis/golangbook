package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s []byte = make([]byte, 0)

	s = strconv.AppendBool(s, true)
	fmt.Println(string(s)) // true: true를 문자열로 변환하여 "true"

	s = strconv.AppendFloat(s, 1.3, 'f', -1, 32) // dst, f, fmt, prec, bitSize
	fmt.Println(string(s)) // true1.3: 1.3을 문자열로 변환하여 "1.3", 슬라이스 뒤에 붙여서 true1.3

	s = strconv.AppendInt(s, -10, 10)
	fmt.Println(string(s)) // true1.3-10: -10을 10진수로된 문자열로 변환하여 "-10", 
                               // 슬라이스 뒤에 붙여서 true1.3-10

	s = strconv.AppendUint(s, 32, 16)
	fmt.Println(string(s)) // true1.3-1020: 32를 16진수로된 문자열로 변환하여 "20", 
                               // 슬라이스 뒤에 붙여서 true1.3-1020
}
