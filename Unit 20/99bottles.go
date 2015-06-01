package main

import "fmt"

func main() {
	for bottles := 99; bottles >= 0; bottles-- {
		switch {
		case bottles > 1:
			fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", bottles, bottles)
			s := "bottles"
			if bottles-1 == 1 {
				s = "bottle"
			}
			fmt.Printf("Take one down, pass it around, %d %s of beer on the wall.\n", bottles-1, s)
		case bottles == 1:
			fmt.Printf("1 bottle of beer on the wall, 1 bottle of beer.\n")
			fmt.Printf("Take one down, pass it around, No more bottles of beer on the wall\n")
		default:
			fmt.Printf("No more bottles of beer on the wall, no more bottles of beer.\n")
			fmt.Printf("Go to the store and buy some more, 99 bottles of beer on the wall.\n")
		}
	}
}
