package main

import "fmt"

func main() {
	fmt.Printf("%3d\n", 1)      //   1
	fmt.Printf("%3d\n", 12345)  // 12345
	fmt.Printf("%03d\n", 1)     // 001
	fmt.Printf("%03d\n", 12345) // 12345

	fmt.Printf("%+d, %d\n", 1, -1) // +1

	fmt.Printf("%9f\n", 1.1234567)   //  1.123457
	fmt.Printf("%09f\n", 1.1234567)  // 01.123457
	fmt.Printf("%.2f\n", 1.1234567)  // 1.12
	fmt.Printf("%9.2f\n", 1234567.0) // 1234567.00
}
