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

	fmt.Printf("정수: %d\n", num1)     // 정수: 10
	fmt.Printf("실수: %f\n", num2)     // 실수: 3.2
	fmt.Printf("복소수: %f\n", num3)   // 복소수: (2.500000+8.100000i)
	fmt.Printf("문자열: %s\n", s)      // 문자열: Hello, world!
	fmt.Printf("불: %t\n", b)          // 불: true
	fmt.Printf("슬라이스: %v\n", a)    // 슬라이스: [1 2 3]
	fmt.Printf("맵: %v\n", m)          // 맵: map[Hello:1]
	fmt.Printf("포인터: %p\n", p)      // 포인터: 0xc0820062d0 (메모리 주소)
	fmt.Printf("구조체: %v\n", data)   // 구조체: {1 2}
	fmt.Printf("인터페이스: %v\n", i)  // 인터페이스: 1

	fmt.Printf("%d %f %s\n", num1, num2, s) // 10 3.200000 Hello, world!
}
