package main

import "fmt"

func main() {
	i := 7

	switch {               // case에 조건식을 지정했으므로 판단할 변수는 생략
	case i >= 5 && i < 10: // i가 5보다 크거나 같으면서 10보다 작을 때
		fmt.Println("5 이상, 10 미만")
	case i >= 0 && i < 5:  // i가 0보다 크거나 같으면서 5보다 작을 때
		fmt.Println("0 이상, 5 미만")
	}
}