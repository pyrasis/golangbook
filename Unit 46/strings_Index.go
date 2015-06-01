package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Index("Hello, world!", "He"))  // 0: He가 맨 처음에 있으므로 0
	fmt.Println(strings.Index("Hello, world!", "wor")) // 7: wor가 8번째에 있으므로 7
	fmt.Println(strings.Index("Hello, world!", "ow"))  // -1: ow는 없으므로 -1

	fmt.Println(strings.IndexAny("Hello, world!", "eo")) // 1: e가 2번째에 있으므로 1
	fmt.Println(strings.IndexAny("Hello, world!", "f"))  // -1: f는 없으므로 -1

	var c byte
	c = 'd'
	fmt.Println(strings.IndexByte("Hello, world!", c)) // 11: d가 12번째에 있으므로 11
	c = 'f'
	fmt.Println(strings.IndexByte("Hello, world!", c)) // -1: f는 없으므로 -1

	var r rune
	r = '언'
	fmt.Println(strings.IndexRune("고 언어", r)) // 4: "언"이 시작되는 인덱스가 4

	f := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r) // r이 한글 유니코드이면 true를 리턴
	}
	fmt.Println(strings.IndexFunc("Go 언어", f))  // 3: 한글이 4번째부터 시작하므로 3
	fmt.Println(strings.IndexFunc("Go Language", f)) // -1: 한글이 없으므로 -1

	fmt.Println(strings.LastIndex("Hello Hello Hello, world!", "Hello")) 
                                                                // 12: 마지막 Hello가 13번째에 있으므로 12

	fmt.Println(strings.LastIndexAny("Hello, world", "ol")) // 10: 마지막 l이 11번째에 있으므로 10

	fmt.Println(strings.LastIndexFunc("Go 언어 안녕", f)) // 13: 마지막 한글인 '녕'이 시작되는 인덱스가 13
}
