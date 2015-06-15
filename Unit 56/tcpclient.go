package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8000") // TCP 프로토콜, 127.0.0.1:8000 서버에 연결
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close() // main 함수가 끝나기 직전에 TCP 연결을 닫음

	go func(c net.Conn) {
		data := make([]byte, 4096) // 4096 크기의 바이트 슬라이스 생성

		for {
			n, err := c.Read(data) // 서버에서 받은 데이터를 읽음
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(data[:n])) // 데이터 출력

			time.Sleep(1 * time.Second)
		}
	}(client)

	go func(c net.Conn) {
		i := 0
		for {
			s := "Hello " + strconv.Itoa(i)

			_, err := c.Write([]byte(s)) // 서버로 데이터를 보냄
			if err != nil {
				fmt.Println(err)
				return
			}

			i++
			time.Sleep(1 * time.Second)
		}
	}(client)

	fmt.Scanln()
}
