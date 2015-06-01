package main

import "fmt"

func main() {
Loop: // Loop 레이블을 지정
	for i := 0; i < 3; i++ {           // 반복문 1
		for j := 0; j < 3; j++ {   // 반복문 2
			if j == 2 {        // j가 2일 때
				break Loop // 중첩된 반복문을 빠져나옴
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")
}
