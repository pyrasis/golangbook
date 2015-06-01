package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var b1 []byte = []byte("안녕하세요")
	fmt.Println(utf8.Valid(b1)) // true: "안녕하세요"는 UTF-8이 맞으므로 true
	var b2 []byte = []byte{0xff, 0xf1, 0xc1}
	fmt.Println(utf8.Valid(b2)) // false: 0xff 0xf1 0xc1은 UTF-8이 아니므로 false

	var r1 rune = '한'
	fmt.Println(utf8.ValidRune(r1)) // true: '한'은 UTF-8이 맞으므로 true
	var r2 rune = 0x11111111
	fmt.Println(utf8.ValidRune(r2)) // false: 0x11111111은 UTF-8이 아니므로 false

	var s1 string = "한글"
	fmt.Println(utf8.ValidString(s1)) // true: "한글"은 UTF-8이 맞으므로 true
	var s2 string = string([]byte{0xff, 0xf1, 0xc1})
	fmt.Println(utf8.ValidString(s2)) // false: 0xff 0xf1 0xc1은 UTF-8이 아니므로 false
}
