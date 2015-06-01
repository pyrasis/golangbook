package main

import "fmt"

func main() {
	var b []byte
	var err error

	if b, err := ioutil.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b)
	}
}