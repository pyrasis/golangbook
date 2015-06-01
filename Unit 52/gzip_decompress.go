package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("hello.txt.gz") // 압축 파일 열기
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	r, err := gzip.NewReader(file) // file로 io.Reader 인터페이스를 따르는 
                                       // 압축 해제 인스턴스 r 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Close() // main 함수가 끝나기 직전에 압축 인스턴스를 닫음

	b, err := ioutil.ReadAll(r) // 압축 해제 인스턴스 r을 읽으면 
                                    // 압축 해제된 데이터가 리턴됨 
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b)) // 문자열로 변환하여 출력
}
