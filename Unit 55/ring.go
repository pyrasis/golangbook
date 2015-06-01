package main

import (
	"container/ring"
	"fmt"
)

func main() {
	data := []string{"Maria", "John", "Andrew", "James"}

	r := ring.New(len(data))       // 노드의 개수를 지정하여 링 생성
	for i := 0; i < r.Len(); i++ { // 링 노드 개수만큼 반복
		r.Value = data[i]      // 링의 노드에 값 넣기
		r = r.Next()           // 다음 노드로  이동
	}

	r.Do(func(x interface{}) { // 링의 모든 노드 순회
		fmt.Println(x)
	})

	fmt.Println("Move forward :")
	r = r.Move(1) // 링을 시계 방향으로 1노드 만큼 회전

	fmt.Println("Curr : ", r.Value)        // Curr :  John
	fmt.Println("Prev : ", r.Prev().Value) // Prev :  Maria
	fmt.Println("Next : ", r.Next().Value) // Next :  Andrew
}
