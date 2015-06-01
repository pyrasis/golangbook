package main

import "fmt"

func main() {
	i := 3

	switch i {                    // 값을 판단할 변수 설정
	case 4:                       // 각 조건에 일치하는
		fmt.Println("4 이상") // 코드를 실행합니다.
		fallthrough
	case 3:                       // 3과 변수의 값이 일치하므로
		fmt.Println("3 이상") // 이 부분을 실행
		fallthrough           // fallthrough를 사용했으므로 아래 case를 모두 실행
	case 2:
		fmt.Println("2 이상") // 실행
		fallthrough
	case 1:
		fmt.Println("1 이상") // 실행
		fallthrough
	case 0:
		fmt.Println("0 이상") // 실행, 마지막 case에는 fallthrough를 사용할 수 없음
	}
}