package main

import (
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

	n := 0
	s := "안녕하세요"
	n, err = file.Write([]byte(s)) // s를 []byte 바이트 슬라이스로 변환, s를 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 저장 완료")

	fi, err := file.Stat() // 파일 정보 가져오기
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // 파일 크기만큼 바이트 슬라이스 생성

	file.Seek(0, os.SEEK_SET) // Write 함수로 인해 파일 읽기/쓰기 위치가 
                                  // 이동했으므로 file.Seek 함수로 읽기/쓰기 위치를 
                                  // 파일의 맨 처음(0)으로 이동
	n, err = file.Read(data)  // 파일의 내용을 읽어서 바이트 슬라이스에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "바이트 읽기 완료")
	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}
