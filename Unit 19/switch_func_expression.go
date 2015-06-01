package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())      // 현재 시간으로 Seed 값 설정
	switch i := rand.Intn(10); {          // rand.Intn 함수를 실행한 뒤 i에 대입
	case i >= 3 && i < 6:                 // i가 3보다 크거나 같으면서 6보다 작을 때
		fmt.Println("3 이상, 6 미만") // 코드 실행
	case i == 9:                          // i가 9일 때
		fmt.Println("9")              // 코드 실행
	default:                              // 모든 case에 해당하지 않을 때
		fmt.Println(i)                // 코드 실행
	}
}