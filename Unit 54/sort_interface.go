package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score float32
}

type Students []Student

func (s Students) Len() int {
	return len(s) // 데이터의 길이를 구함. 슬라이스이므로 len 함수를 사용
}

func (s Students) Less(i, j int) bool {
	return s[i].name < s[j].name // 학생 이름순으로 정렬
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i] // 두 데이터의 위치를 바꿈
}

type ByScore struct { // 점수순 정렬을 위해 구조체 정의
	Students
}

func (s ByScore) Less(i, j int) bool {
	return s.Students[i].score < s.Students[j].score // 학생 점수순으로 정렬
}

func main() {
	s := Students{
		{"Maria", 89.3},
		{"Andrew", 72.6},
		{"John", 93.1},
	}

	sort.Sort(s) // 이름을 기준으로 오름차순 정렬
	fmt.Println(s)

	sort.Sort(sort.Reverse(ByScore{s})) // 점수를 기존으로 내림차순 정렬
	fmt.Println(s)
}
