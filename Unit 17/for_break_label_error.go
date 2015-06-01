package main

import "fmt"

func main() {
Loop:
	fmt.Println("begin for loop") // 들어가면 안 되는 코드
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop // 컴파일 에러. invalid break label Loop
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")
}
