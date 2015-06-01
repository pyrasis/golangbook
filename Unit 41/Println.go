package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 float32 = 3.2
	var num3 complex64 = 2.5 + 8.1i
	var s string = "Hello, world!"
	var b bool = true
	var a []int = []int{1, 2, 3}
	var m map[string]int = map[string]int{"Hello": 1}
	var p *int = new(int)
	type Data struct{ a, b int }
	var data Data = Data{1, 2}
	var i interface{} = 1

	fmt.Println(num1) // 10: 정수 출력
	fmt.Println(num2) // 3.2: 실수 출력
	fmt.Println(num3) // (2.5+8.1i): 복소수 출력
	fmt.Println(s)    // Hello, world!: 문자열 출력
	fmt.Println(b)    // true: 불 출력
	fmt.Println(a)    // [1 2 3]: 슬라이스 출력
	fmt.Println(m)    // map[Hello:1]: 맵 출력
	fmt.Println(p)    // 0xc0820062d0: 포인터(메모리 주소) 출력
	fmt.Println(data) // {1 2}: 구조체 출력
	fmt.Println(i)    // 1: 인터페이스 출력

	fmt.Println(num1, num2, num3, s, b) // 10 3.2 (2.5+8.1i) Hello, world! true
	fmt.Println(p, a, m)                // 0xc0820062d0 [1 2 3] map[Hello:1]
	fmt.Println(data, i)                // {1 2} 1
}
