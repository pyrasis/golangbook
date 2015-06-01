package main

import (
	"fmt"
	"regexp"
)

func main() {
	var matched, _ = regexp.MatchString("^Hello", "Hello, world!")
	fmt.Println(matched) // true: Hello, world!에서 Hello가 맨 앞에 오므로 true

	matched, _ = regexp.MatchString("^Hello", "Go Hello, world!")
	fmt.Println(matched) // false: Go Hello, world!에서 Hello가 맨 앞에 오지 않으므로 false

	matched, _ = regexp.MatchString("world!$", " Hello, world!")
	fmt.Println(matched) // true: Hello, world!에서 world!가 맨 뒤에 오므로 true

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "Hello.SetValue(20)")
	fmt.Println(matched) // true: Hello.SetValue(20)는 .영문(숫자)이므로 true

	matched, _ = regexp.MatchString("\\.[a-zA-Z]+\\([0-9]+\\)$", "Hello.SetValue(20).x")
	fmt.Println(matched) // false: Hello.SetValue(20).x는 .영문(숫자)가 아니므로 false
}
