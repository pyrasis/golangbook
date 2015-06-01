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
	data := make([]Article, 1) // 값을 저장할 구조체 슬라이스 생성

	data[0].Id = 1
	data[0].Title = "Hello, world!"
	data[0].Author.Name = "Maria"
	data[0].Author.Email = "maria@example.com"
	data[0].Content = "Hello~"
	data[0].Recommends = []string{"John", "Andrew"}
	data[0].Comments = make([]Comment, 1)
	data[0].Comments[0].Id = 1
	data[0].Comments[0].Author.Name = "Andrew"
	data[0].Comments[0].Author.Email = "andrew@hello.com"
	data[0].Comments[0].Content = "Hello Maria"

	doc, _ := json.Marshal(data) // data를 JSON 문서로 변환

	fmt.Println(string(doc)) // [{"Id":1,"Title":"Hello, world!","Au... (생략)
}
