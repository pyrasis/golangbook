package main

import (
	"fmt"
	"strconv"
)

//                      ↓ 빈 인터페이스를 사용하여 모든 타입을 받음
func formatString(arg interface{}) string {
	//       ↓ 인터페이스에 저장된 타입에 따라 case 실행
	switch arg.(type) {
	case int:                      // arg가 int형이라면
		i := arg.(int)         // int형으로 값을 가져옴
		return strconv.Itoa(i) // strconv.Itoa 함수를 사용하여 i를 문자열로 변환
	case float32:                  // arg가 float32형이라면
		f := arg.(float32)     // float32형으로 값을 가져옴
		return strconv.FormatFloat(float64(f), 'f', -1, 32) 
                                       // strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case float64:                  // arg가 float64형이라면
		f := arg.(float64)     // float64형으로 값을 가져옴
		return strconv.FormatFloat(f, 'f', -1, 64)
                                       // strconv.FormatFloat 함수를 사용하여 f를 문자열로 변환
	case string:                   // arg가 string이라면
		s := arg.(string)      // string으로 값을 가져옴
		return s               // string이므로 그대로 리턴
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello, world!"))
}
