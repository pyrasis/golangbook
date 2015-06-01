package main

import (
	"fmt"
	"os"
)

func main() {
	file1, _ := os.Create("hello1.txt")        // hello1.txt 파일 생성
	defer file1.Close()                        // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprint(file1, 1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만든 뒤 파일에 저장

	file2, _ := os.Create("hello2.txt")          // hello2.txt 파일 생성
	defer file2.Close()                          // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprintln(file2, 1, 1.1, "Hello, world!") // 값을 그대로 문자열로 만든 뒤 
	                                             // 문자열 끝에 개행 문자(\n)를 붙이고 파일에 저장

	file3, _ := os.Create("hello3.txt")                     // hello3.txt 파일 생성
	defer file3.Close()                                     // main 함수가 끝나기 직전에 파일을 닫음
	fmt.Fprintf(file3, "%d,%f,%s", 1, 1.1, "Hello, world!") // 형식을 지정하여 파일에 저장
}
