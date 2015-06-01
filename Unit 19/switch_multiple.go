package main

import "fmt"

func main() {
	i := 3

	switch i {
	case 2, 4, 6: // i가 2, 4, 6일 때
		fmt.Println("짝수")
	case 1, 3, 5: // i가 1, 3, 5일 때
		fmt.Println("홀수")
	}
}