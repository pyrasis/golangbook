package main

import "fmt"

func main() {
Loop:
	for i := 0; i < 3; i++ {              // 반복문 1
		for j := 0; j < 3; j++ {      // 반복문 2
			if j == 2 {           // j가 2일 때
				continue Loop // 아래 부분 코드를 실행하지 않고 반복문 1부터 이어서 실행
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")
}
