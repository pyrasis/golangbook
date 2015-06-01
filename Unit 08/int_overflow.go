package main

import "fmt"
import "math"

func main() {
	var num1 uint8 = math.MaxUint8 + 1   // 오버플로우 컴파일 에러 발생
	var num2 uint16 = math.MaxUint16 + 1 // 오버플로우 컴파일 에러 발생
	var num3 uint32 = math.MaxUint32 + 1 // 오버플로우 컴파일 에러 발생
	var num4 uint64 = math.MaxUint64 + 1 // 오버플로우 컴파일 에러 발생
}
