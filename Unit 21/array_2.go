package main

import "fmt"

func main() {
	var a [5]int = [5]int{32, 29, 78, 16, 81} // int형이며 길이가 5인 배열을 선언하고 초기화
	var b = [5]int{32, 29, 78, 16, 81}        // 배열을 선언할 때 자료형과 길이 생략
	c := [5]int{32, 29, 78, 16, 81}           // 배열을 선언할 때 var 키워드, 자료형과 길이 생략

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}