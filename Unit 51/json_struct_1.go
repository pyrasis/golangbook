package main

import (
	"encoding/json"
	"fmt"
)

type Author struct {
	Name  string
	Email string
}

type Comment struct {
	Id      uint64
	Author  Author // Author 구조체
	Content string
}

type Article struct {
	Id         uint64
	Title      string
	Author     Author    // Author 구조체
	Content    string
	Recommends []string  // 문자열 배열
	Comments   []Comment // Comment 구조체 배열
}

func main() {
	doc := `
	[{
		"Id": 1,
		"Title": "Hello, world!",
		"Author": {
			"Name": "Maria",
			"Email": "maria@example.com"
		},
		"Content": "Hello~",
		"Recommends": [
			"John",
			"Andrew"
		],
		"Comments": [{
			"id": 1,
			"Author": {
				"Name": "Andrew",
				"Email": "andrew@hello.com"
			},
			"Content": "Hello Maria"
		}]
	}]
	`

	var data []Article // JSON 문서의 데이터를 저장할 구조체 슬라이스 선언

	json.Unmarshal([]byte(doc), &data) // doc의 내용을 변환하여 data에 저장

	fmt.Println(data) // [{1 Hello, world! {Maria maria@exa... (생략)
}
