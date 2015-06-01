package main

import "fmt"

func main() {
	var a map[string]int = make(map[string]int) // make 함수로 키는 string 값은 int인 맵에 공간 할당
	var b = make(map[string]int)                // 맵을 선언할 때 map 키워드와 자료형 생략
	c := make(map[string]int)                   // 맵을 선언할 때 var, map 키워드와 자료형 생략

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
