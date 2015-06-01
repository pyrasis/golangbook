package main

import "fmt"
import "runtime"

func main() {
	fmt.Println("CPU Count : ", runtime.NumCPU())
}
