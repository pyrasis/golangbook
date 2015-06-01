package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")              // 바꿀 문자열을 지정
	fmt.Println(r.Replace("<div><span>Hello, world!</span></div>")) // HTML에서 < >를 &lt; &gt;로 바꿈
}
