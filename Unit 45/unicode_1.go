package main

import (
	"fmt"
	"unicode"
)

func main() {
	var r1 rune = '한'
	fmt.Println(unicode.Is(unicode.Hangul, r1)) // true: r1은 한글이므로 true
	fmt.Println(unicode.Is(unicode.Latin, r1))  // false: r1은 라틴 문자가
	                                            // 아니므로 false

	var r2 rune = '漢'
	fmt.Println(unicode.Is(unicode.Han, r2))    // true: r2는 한자이므로 true
	fmt.Println(unicode.Is(unicode.Hangul, r2)) // false: r2는 한글이 아니므로 false

	var r3 rune = 'a'
	fmt.Println(unicode.Is(unicode.Latin, r3))  // true: r3은 라틴 문자이므로 true
	fmt.Println(unicode.Is(unicode.Hangul, r3)) // false: r3은 한글이 아니므로 false

	fmt.Println(unicode.In(r1, unicode.Latin, unicode.Han, unicode.Hangul)) // true: r1은 한글이므로 true
}
