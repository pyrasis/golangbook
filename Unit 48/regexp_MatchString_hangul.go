package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("[가-힣]+", "안녕하세요")
	fmt.Println(matched) // true

	matched, _ = regexp.MatchString("홍[가-힣]+동", "홍길동")
	fmt.Println(matched) // true

	matched, _ = regexp.MatchString("이재[홍-훈]", "이재홍")
	fmt.Println(matched) // true
}
