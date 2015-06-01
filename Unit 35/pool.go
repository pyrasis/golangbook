package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Data struct { // Data 구조체 정의
	tag    string  // 풀 태그
	buffer []int   // 데이터 저장용 슬라이스
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	pool := sync.Pool{                            // 풀 할당
		New: func() interface{} {             // Get 함수를 사용했을 때 호출될 함수 정의
			data := new(Data)             // 새 메모리 할당
			data.tag = "new"              // 태그 설정
			data.buffer = make([]int, 10) // 슬라이스 공간 할당
			return data                   // 할당한 메모리(객체) 리턴
		},
	}

	for i := 0; i < 10; i++ {
		go func() {                                          // 고루틴 10개 생성
			data := pool.Get().(*Data)                   // 풀에서 *Data 타입으로 데이터를 가져옴
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100)  // 슬라이스에 랜덤 값 저장
			}
			fmt.Println(data)                            // data 내용 출력
			data.tag = "used"                            // 객체가 사용되었다는 태그 설정
			pool.Put(data)                               // 풀에 객체를 보관
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {                               // 고루틴 10개 생성
			data := pool.Get().(*Data)        // 풀에서 *Data 타입으로 데이터를 가져옴
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n    // 슬라이스에 짝수 저장
				n += 2
			}
			fmt.Println(data)                 // data 내용 출력
			data.tag = "used"                 // 객체가 사용되었다는 태그 설정
			pool.Put(data)                    // 풀에 객체 보관
		}()
	}

	fmt.Scanln()
}
