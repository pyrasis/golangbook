package main

import "fmt"

func main() {
	a := [5]int{32, 29, 78, 16, 81} // 배열을 생성하면서 값을 초기화

	b := [5]int{
		32,
		29,
		78,
		16,
		81, // 여러 줄로 나열할 때는 마지막 요소에 콤마를 붙임
	}

	c := [...]int{32, 29, 78, 16, 81} // 초기화할 요소가 5개이며 ...을 사용했으므로 
	                                  // 배열 크기는 5로 설정됨

	d := [...]string{"Maria", "Andrew", "Jonh"} // 초기화할 요소가 3개이며 ...을 사용했으므로 
	                                            // 배열 크기는 3으로 설정됨

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}