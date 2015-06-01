package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	score float32
}

type By func(s1, s2 *Student) bool // 각 상황별 정렬 함수를 저장할 타입

func (by By) Sort(students []Student) { // By 함수 타입에 메서드를 붙임
	sorter := &studentSorter{       // 데이터와 정렬 함수로 sorter 초기화
		students: students,
		by:       by,
	}
	sort.Sort(sorter) // 정렬
}

type studentSorter struct {
	students []Student                  // 데이터
	by       func(s1, s2 *Student) bool // 각 상황별 정렬 함수
}

func (s *studentSorter) Len() int {
	return len(s.students) // 슬라이스의 길이를 구함
}

func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.students[i], &s.students[j]) // by 함수로 대소관계 판단
}

func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i] // 데이터 위치를 바꿈
}

func main() {
	s := []Student{
		{"Maria", 89.3},
		{"Andrew", 72.6},
		{"John", 93.1},
	}

	name := func(p1, p2 *Student) bool { // 이름 오름차순 정렬 함수
		return p1.name < p2.name
	}
	score := func(p1, p2 *Student) bool { // 점수 오름차순 정렬 함수
		return p1.score < p2.score
	}
	reverseScore := func(p1, p2 *Student) bool { // 점수 내림차순 정렬 함수
		return !score(p1, p2)
	}

	By(name).Sort(s) // name 함수를 사용하여 이름 오름차순 정렬
	fmt.Println(s)

	By(reverseScore).Sort(s) // reverseScore 함수를 사용하여 점수 내림차순 정렬
	fmt.Println(s)
}
