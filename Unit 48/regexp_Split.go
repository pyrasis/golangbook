package main

import (
	"fmt"
	"regexp"
)

func main() {
	re1, _ := regexp.Compile("\\.([a-z]+)\\.")
	s1 := re1.Split("mail.hello.net www.example.com ftp.example.org", -1) 
                        // 정규표현식에 해당하는 문자열일 기준으로 모든 문자열을 쪼갬(-1)
	fmt.Println(s1) // [mail net www com ftp org]: ".영문."을 기준으로 쪼갠 결과

	re2, _ := regexp.Compile("\\.([a-z]+)\\.")
	s2 := re2.Split("mail.hello.net www.example.com ftp.example.org", 2) 
                        // 정규표현식에 해당하는 문자열일 기준으로 문자열을 쪼갬
	fmt.Println(s2) // [mail net www.example.com ftp.example.org]: 마지막 문자열을 
                        // 제외한 1개 문자열을 쪼갬

	re3, _ := regexp.Compile("\\.([a-z]+)\\.")
	s3 := re3.Split("mail.hello.net www.example.com ftp.example.org", 3)
                        // 정규표현식에 해당하는 문자열일 기준으로 문자열을 쪼갬
	fmt.Println(s3) // [mail net www com ftp.example.org]: 마지막 문자열을 제외한 
                        // 2개 문자열을 쪼갬

	re4, _ := regexp.Compile("\\.([a-z]+)\\.")
	s4 := re4.Split("mail.hello.net www.example.com ftp.example.org", 0) 
                        // 정규표현식에 해당하는 문자열일 기준으로 문자열을 쪼갬
	fmt.Println(s4) // []: 아무 문자열도 쪼개지 않음
}
