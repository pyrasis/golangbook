package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 2 {  // i가 2일 때
			continue // 아래 부분 코드를 실행하지 않고 넘어감
		}

		fmt.Println(i)
	}
}
