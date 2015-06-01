package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int // 힙을 int 슬라이스로 정의

func (h MinHeap) Len() int {
	return len(h) // 슬라이스의 길이를 구함
}

func (h MinHeap) Less(i, j int) bool {
	r := h[i] < h[j] // 대소관계 판단
	fmt.Printf("Less %d < %d %t\n", h[i], h[j], r)
	return r
}

func (h MinHeap) Swap(i, j int) {
	fmt.Printf("Swap %d %d\n", h[i], h[j])
	h[i], h[j] = h[j], h[i] // 값의 위치를 바꿈
}

func (h *MinHeap) Push(x interface{}) {
	fmt.Println("Push", x)
	*h = append(*h, x.(int)) // 맨 마지막에 값 추가
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]     // 슬라이스의 맨 마지막 값을 가져옴
	*h = old[0 : n-1] // 맨 마지막 값을 제외한 슬라이스를 다시 저장
	return x
}

func main() {
	data := new(MinHeap) // 힙 생성

	heap.Init(data)    // 힙 초기화
	heap.Push(data, 5) // 힙에 데이터 추가
	heap.Push(data, 2)
	heap.Push(data, 7)
	heap.Push(data, 3)

	fmt.Println(data, "최솟값 : ", (*data)[0])
}
