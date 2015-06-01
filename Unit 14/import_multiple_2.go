package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU Count : ", runtime.NumCPU())
}
