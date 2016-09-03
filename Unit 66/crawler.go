package main

import (
	"fmt"
	"golang.org/x/net/html" // HTML 파싱 해키지
	"log"
	"net/http"
	"runtime"
	"sync"
)

var fetched = struct {
	m map[string]error  // 중복 검사를 위한 URL과 에러 값 저장
	sync.Mutex
}{m: make(map[string]error)} // 변수를 선언하면서 이름이 없는 구조체를 정의하고 초깃값을 생성하여 대입

func fetch(url string) (*html.Node, error) {
	res, err := http.Get(url) // URL에서 HTML 데이터를 가져옴
	if err != nil {
		log.Println(err)
		return nil, err
	}

	doc, err := html.Parse(res.Body) // res.Body를 넣으면 파싱된 데이터가 리턴됨
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return doc, nil
}

func parseFollowing(doc *html.Node) []string {
	var urls = make([]string, 0)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" { // img 태그
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "avatar float-left" { // class가 avatar float-left인 요소
					for _, a := range n.Attr {
						if a.Key == "alt" {
							fmt.Println(a.Val) // 사용자 이름 출력
							break
						}
					}
				}

				if a.Key == "class" && a.Val == "avatar" { // class가 avatar인 요소
					user := n.Parent.Attr[0].Val // 부모 요소의 첫 번째 속성(href)

					// 사용자 이름으로 팔로잉 URL 조합
					urls = append(urls, "https://github.com"+user+"/following") 
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c) // 재귀호출로 자식과 형제를 모두 탐색
		}
	}
	f(doc)

	return urls
}

func crawl(url string) {
	fetched.Lock()                   // 맵은 뮤텍스로 보호
	if _, ok := fetched.m[url]; ok { // URL 중복 처리 여부를 검사
		fetched.Unlock()
		return
	}
	fetched.Unlock()

	doc, err := fetch(url) // URL에서 파싱된 HTML 데이터를 가져옴

	fetched.Lock()
	fetched.m[url] = err // 가져온 URL은 맵에 URL과 에러 값 저장
	fetched.Unlock()

	urls := parseFollowing(doc) // 사용자 정보 출력, 팔로잉 URL을 구함

	done := make(chan bool)
	for _, u := range urls {      // URL 개수만큼
		go func(url string) { // 고루틴 생성
			crawl(url)    // 재귀호출
			done <- true
		}(u)
	}

	for i := 0; i < len(urls); i++ {
		<-done // 고루틴이 모두 실행될 때까지 대기
	}
}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs) // 모든 CPU를 사용하도록 설정

	crawl("https://github.com/pyrasis/following") // 크롤링 시작
}
