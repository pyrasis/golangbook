package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux() // HTTP 요청 멀티플렉서 생성

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!")) // 웹 브라우저에 응답
	})

	http.ListenAndServe(":80", mux) // 80번 포트에 웹 서버를 실행하고, mux를 이용하여 HTTP 요청을 처리
}
