package main

import "fmt"
import "unicode/utf8"

func main() {
	var s1 string = "한글"
	fmt.Println(utf8.RuneCountInString(s1)) // 2: 문자열의 실제 길이를 구함
}
