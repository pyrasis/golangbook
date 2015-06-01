package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.IsGraphic('1'))  // true: 1은 화면에 표시되는 숫자이므로 true
	fmt.Println(unicode.IsGraphic('a'))  // true: a는 화면에 표시되는 문자이므로 true
	fmt.Println(unicode.IsGraphic('한')) // true: '한'은 화면에 표시되는 문자이므로 true
	fmt.Println(unicode.IsGraphic('漢')) // true: '漢'은 화면에 표시되는 문자이므로 true
	fmt.Println(unicode.IsGraphic('\n')) // false: \n 화면에 표시되는 문자가 아니므로 false

	fmt.Println(unicode.IsLetter('a')) // true: a는 문자이므로 true
	fmt.Println(unicode.IsLetter('1')) // false: 1은 문자가 아니므로 false

	fmt.Println(unicode.IsDigit('1'))     // true: 1은 숫자이므로 true
	fmt.Println(unicode.IsControl('\n'))  // true: \n은 제어 문자이므로 true
	fmt.Println(unicode.IsMark('\u17c9')) // true: \u17c9는 마크이므로 true

	fmt.Println(unicode.IsPrint('1')) // true: 1은 Go 언어에서 표시할 수 있으므로 true
	fmt.Println(unicode.IsPunct('.')) // true: .은 문장 부호이므로 true

	fmt.Println(unicode.IsSpace(' '))   // true: ' '는 공백이므로 true
	fmt.Println(unicode.IsSymbol('♥')) // true: ♥는 심볼이므로 true

	fmt.Println(unicode.IsUpper('A')) // true: A는 대문자이므로 true
	fmt.Println(unicode.IsLower('a')) // true: a는 소문자이므로 true
}
