package main

import (
	"bufio"
	"fmt"
	"os"
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

	r := bufio.NewReader(file) // file로 io.Reader 인터페이스를 따르는 읽기 인스턴스 r 생성
	w := bufio.NewWriter(file) // file로 io.Writer 인터페이스를 따르는 쓰기 인스턴스 w 생성

	rw := bufio.NewReadWriter(r, w) // r, w를 사용하여 io.ReadWriter 인터페이스를 따르는 
                                        // 읽기/쓰기 인스턴스 생성
	rw.WriteString("Hello, world!") // 읽기/쓰기 인스턴스로 버퍼에 Hello, world! 쓰기
	rw.Flush()                      // 버퍼의 내용을 파일에 저장

	file.Seek(0, os.SEEK_SET)   // 파일 읽기 위치를 파일의 맨 처음(0)으로 이동
	data, _, _ := rw.ReadLine() // 파일에서 문자열 한 줄을 읽어서 data에 저장
	fmt.Println(string(data))   // 문자열로 변환하여 data의 내용 출력
}
