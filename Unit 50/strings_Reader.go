package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC, // 파일이 없으면 생성, 
                                                  // 읽기/쓰기, 파일을 연 뒤 내용 삭제
		os.FileMode(0644))                // 파일 권한은 644
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // main 함수가 끝나기 직전에 파일을 닫음

	s := "Hello, world!"
	r := strings.NewReader(s) // 문자열로 io.Reader 인터페이스를 따르는 
                                  // 읽기 인스턴스 r 생성

	w := bufio.NewWriter(file) // io.Writer 인터페이스를 따르는 file로 
                                   // io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성
	w.ReadFrom(r)              // 읽기 인스턴스 r에서 데이터를 읽어서 w의 버퍼에 저장
	w.Flush()                  // 버퍼의 내용을 파일에 저장
}
