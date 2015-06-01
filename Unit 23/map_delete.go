package main

import "fmt"

func main() {
	a := map[string]int{"Hello": 10, "world": 20}

	delete(a, "world") // 맵 a에서 world 키 삭제

	fmt.Println(a) // map[Hello:10]
}
