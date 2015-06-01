package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Hello, world!", "wo"))  // true
	fmt.Println(strings.Contains("Hello, world!", "w o")) // false
	fmt.Println(strings.Contains("Hello, world!", "ow"))  // false

	fmt.Println(strings.ContainsAny("Hello, world!", "wo"))  // true
	fmt.Println(strings.ContainsAny("Hello, world!", "w o")) // true
	fmt.Println(strings.ContainsAny("Hello, world!", "ow"))  // true

	fmt.Println(strings.Count("Hello Helium", "He")) // 2

	var r rune
	r = '하'
	fmt.Println(strings.ContainsRune("안녕하세요", r)) // true

	fmt.Println(strings.HasPrefix("Hello, world!", "He"))   // true
	fmt.Println(strings.HasSuffix("Hello, world!", "rld!")) // true
}
