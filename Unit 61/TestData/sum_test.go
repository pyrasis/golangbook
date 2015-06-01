package calc

import "testing"

type TestData struct { // 테스트 데이터 구조체
	argument1 int
	argument2 int
	result    int
}

var testdata = []TestData{ // 테스트 데이터 슬라이스
	{2, 6, 8},
	{-8, 3, -5},
	{6, -6, 0},
	{0, 0, 0},
}

func TestSum(t *testing.T) {
	for _, d := range testdata {               // 반복문으로 데이터를 꺼낸 뒤
		r := Sum(d.argument1, d.argument2) // Sum 함수를 테스트
		if r != d.result {
			t.Errorf(
				"%d + %d의 결과값이 %d(이)가 아닙니다. r=%d",
				d.argument1,
				d.argument2,
				d.result,
				r,
			)
		}
	}
}
