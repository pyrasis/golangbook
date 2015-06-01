package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, _ := regexp.MatchString("He", "Hello 100")
	fmt.Println(matched) // true: Hello 100에 He가 포함되므로 true

	matched, _ = regexp.MatchString("H.", "Hi")
	fmt.Println(matched) // true: H 뒤에 글자가 하나가 있으므로 true

	matched, _ = regexp.MatchString("[a-zA-Z0-9]+", "Hello 100")
	fmt.Println(matched) // true: Hello 100은 a부터 z까지, A부터 Z까지 0부터 9까지에 포함되므로 true

	matched, _ = regexp.MatchString("[a-zA-Z0-9]*", "")
	fmt.Println(matched) // true: *는 값이 하나도 안나와도 되므로 true

	matched, _ = regexp.MatchString("[0-9]+-[0-9]+", "1-1")
	fmt.Println(matched) // true: 1-1는 [0-9]+와 - 그리고 [0-9]+가 합쳐진 것이므로 true

	matched, _ = regexp.MatchString("[0-9]+-[0-9]+", "1-")
	fmt.Println(matched) // false: 1-는 [0-9]+와 -만 있으므로 false

	matched, _ = regexp.MatchString("[0-9]+/[0-9]*", "1/")
	fmt.Println(matched) // true: 1/는 [0-9]+와 /가 있음. *는 값이 없어도 되므로 true

	matched, _ = regexp.MatchString("[0-9]+\\+[0-9]*", "1+")
	fmt.Println(matched) // true: 1+는 [0-9]+와 +가 있음. *는 값이 없어도 되므로 true

	matched, _ = regexp.MatchString("[^A-Z]+", "hello")
	fmt.Println(matched) // true: hello는 A부터 Z까지에 포함되지 않으므로 true
}
