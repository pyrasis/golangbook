package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.ReplaceAllString("hello example.com", ".net") 
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s1) // hello example.net: 맨 마지막에 오는 ".영문"을 .net으로 바꿈

	re2, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s2 := re2.ReplaceAllString("Hello, world!", "${2}") 
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s2) // world!: ${2}만 있으므로 두 번째로 찾은 문자열 world!만 남김

	re3, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s3 := re3.ReplaceAllString("Hello, world!", "${2} ${1}")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s3) // world! Hello,: ${2} ${1}로 설정했으므로 두 문자열의 위치를 바꿈

	re4, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)") 
                        // 찾은 문자열에 ${first}, ${second}로 이름을 정함
	s4 := re4.ReplaceAllString("Hello, world!", "${second} ${first}")
                        // 정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈
	fmt.Println(s4) // world! Hello,: ${second} ${first}로 설정했으므로 
                        // 두 문자열의 위치를 바꿈

	re5, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
	s5 := re5.ReplaceAllLiteralString("Hello, world!", "${second} ${first}")
                        // 문자열을 바꿀 때 정규표현식 기호를 무시함
	fmt.Println(s5) // ${second} ${first}: 정규표현식 기호를 무시했으므로 
                        // ${second} ${first}가 그대로 표시됨
}
