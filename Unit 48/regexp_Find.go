package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.FindString("hello example.com") // 정규표현식으로 문자열 검색
	fmt.Println(s1)                           // .com: 맨 마지막에 오는 ".영문"이므로 .com

	re2, _ := regexp.Compile("([a-zA-Z]+)(\\.[a-zA-Z]+)$")
	s2 := re2.FindStringSubmatch("hello example.com") // 문자열을 검색한 뒤 찾은 문자열과
                                                          // 괄호로 구분된 하위 항목 리턴
	fmt.Println(s2) // [example.com example .com]: 맨 마지막에 오는 "영문.영문"을 찾고 
                        // 하위 항목 example과 .com을 리턴

	n2 := re2.FindStringSubmatchIndex("example.com") // 문자열을 검색한 뒤 찾은 문자열의
                                                         // 위치와 괄호로 구분한 하위 항목의 위치를 리턴
	fmt.Println(n2) // [0 11 0 7 7 11]

	re3, _ := regexp.Compile("\\.[a-zA-Z.]+")
	s3 := re3.FindAllString("hello.co.kr world hello.net example.com", -1)
                        // 정규표현식에 해당하는 모든 문자열 검색
	fmt.Println(s3) // [.co.kr .net .com]: ".영문."에 해당하는 모든 문자열(-1)

	s3 = re3.FindAllString("hello.co.kr world hello.net example.com", 2)
                        // 정규표현식에 해당하는 모든 문자열 검색
	fmt.Println(s3) // [.co.kr .net]:  ".영문."에 해당하는 문자열 2개 리턴

	n3 := re3.FindAllStringIndex("hello.co.kr world hello.net example.com", -1)
                        // 정규표현식에 해당하는 모든 문자열의 위치를 리턴
	fmt.Println(n3) // [[5 11] [23 27] [35 39]]: ".영문."에 해당하는 모든 문자열의 위치
}
