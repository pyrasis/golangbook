package main

import (
	"flag"
	"fmt"
)

func main() {
	title := flag.String("title", "", "영화 이름")      // 명령줄 옵션을 받은 뒤 문자열로 저장
	runtime := flag.Int("runtime", 0, "상영 시간")      // 명령줄 옵션을 받은 뒤 정수로 저장
	rating := flag.Float64("rating", 0.0, "평점")       // 명령줄 옵션을 받은 뒤 실수로 저장
	release := flag.Bool("release", false, "개봉 여부") // 명령줄 옵션을 받은 뒤 불로 저장

	flag.Parse() // 명령줄 옵션의 내용을 각 자료형별로 분석

	if flag.NFlag() == 0 { // 명령줄 옵션의 개수가 0개이면
		flag.Usage()   // 명령줄 옵션 기본 사용법 출력
		return
	}

	fmt.Printf(
		"영화 이름: %s\n상영 시간: %d분\n평점: %f\n",
		*title, // 포인터이므로 값을 꺼낼 때는 역참조 사용
		*runtime,
		*rating,
	) // 명령줄 옵션으로 받은 값을 출력
	
	if *release == true {
		fmt.Println("개봉 여부: 개봉")
	} else {
		fmt.Println("개봉 여부: 미개봉")
	}
}
