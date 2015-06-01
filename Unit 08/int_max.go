package main

import "fmt"
import "math"

func main() {
	var num1 uint8 = math.MaxUint8
	var num2 uint16 = math.MaxUint16
	var num3 uint32 = math.MaxUint32
	var num4 uint64 = math.MaxUint64

	fmt.Println(num1) // 255
	fmt.Println(num2) // 65535
	fmt.Println(num3) // 4294967295
	fmt.Println(num4) // 18446744073709551615
}

