package main

import "fmt"

func main() {
	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
		bit1, mask1                          // bit1 == 2, mask1 == 1
		_, _                                 // iota == 2를 생략하여 bit2와 mask2 생략
		bit3, mask3                          // bit3 == 8, mask3 == 7
	)

	fmt.Println(bit0)
}
