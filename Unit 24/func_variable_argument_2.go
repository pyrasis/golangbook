package main

import "fmt"

func sum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}

	return total
}

func main() {
	n := []int{1, 2, 3, 4, 5}
	r := sum(n...) // ...를 사용하여 가변인자에 
                        // 슬라이스를 바로 넘겨줌

	fmt.Println(r) // 15
}
