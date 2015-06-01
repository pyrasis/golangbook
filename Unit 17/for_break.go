package main

import "fmt"

func main() {
	i := 0
	for {
		if i > 4 { // i가 4가 되는 순간
			break  // 반복문이 중단됩니다.
		}

		fmt.Println(i)
		i = i + 1 // 변화식에서 조건을 변경합니다.
	}
}
