package main

import "fmt"

func main() {
	var a []int = make([]int, 5) // make 함수로 int형에 길이가 5인 슬라이스에 공간 할당
	var b = make([]int, 5)       // 슬라이스를 선언할 때 자료형과 [] 생략
	c := make([]int, 5)          // 슬라이스를 선언할 때 var 키워드, 자료형과 [] 생략

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}