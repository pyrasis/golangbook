package main

import "fmt"

func main() {
	a := make([]int, 5, 10) // 길이가 5이면 a[0], a[1], a[2], a[3], a[4]가 생성

	fmt.Println(a[4]) // 0: make 함수를 사용하면 슬라이스의 요소는 모두 0으로 초기화
	fmt.Println(a[5]) // 길이를 벗어난 인덱스에 접근했으므로 런타임 에러 발생
	fmt.Println(a[8]) // 길이를 벗어난 인덱스에 접근했으므로 런타임 에러 발생
}