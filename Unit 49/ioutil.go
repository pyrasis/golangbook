package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	s := "Hello, world!"

	err := ioutil.WriteFile("hello.txt", []byte(s), os.FileMode(644)) 
                                                  // s를 []byte 바이트 슬라이스로 변환, s를 hello.txt 파일에 저장
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadFile("hello.txt") // hello.txt의 내용을 읽어서 바이트 슬라이스 리턴
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력
}
