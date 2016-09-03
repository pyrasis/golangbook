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
	m map[string]error   // 중복 검사를 위한 URL과 에러 값 저장
	sync.Mutex
}{m: make(map[string]error)} // 변수를 선언하면서 이름이 없는 구조체를 정의하고 
                             // 초깃값을 생성하여 대입

type result struct { // 결괏값을 저장할 구조체
	url  string  // 가져온 URL
	name string  // 사용자 이름
}

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

func parseFollowing(doc *html.Node, urls chan string) <-chan string {
	name := make(chan string)

	go func() { // 교착 상태가 되지 않도록 고루틴으로 실행
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "img" { // img 태그
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "avatar float-left" { // class가 avatar float-left인 요소
						for _, a := range n.Attr {
							if a.Key == "alt" {
								// 사용자 이름을 구한 뒤 채널에 보냄
								name <- a.Val 
								break
							}
						}
					}

					if a.Key == "class" && a.Val == "avatar" { // class가 avatar인 요소
						user := n.Parent.Attr[0].Val // 부모 요소의 첫 번째 속성(href)

						// 사용자 이름으로 팔로잉 URL을 조합하여 urls 채널에 보냄
						urls <- "https://github.com" + user + "/following" 

						break
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c) // 재귀호출로 자식과 형제를 모두 탐색
			}
		}
		f(doc)
	}()

	return name // 채널을 리턴
}

func crawl(url string, urls chan string, c chan<- result) {
	fetched.Lock()                   // 맵은 뮤텍스로 보호
	if _, ok := fetched.m[url]; ok { // URL 중복 처리 여부를 검사
		fetched.Unlock()
		return
	}
	fetched.Unlock()

	doc, err := fetch(url)      // URL에서 파싱된 HTML 데이터를 가져옴
	if err != nil {             // URL을 가져오지 못했을 때
		go func(u string) { // 교착 상태가 되지 않도록 고루틴을 생성
			urls <- u   // 채널에 URL을 보냄
		}(url)
		return
	}

	fetched.Lock()
	fetched.m[url] = err // 가져온 URL은 맵에 URL과 에러 값 저장
	fetched.Unlock()

	name := <-parseFollowing(doc, urls) // 사용자 정보, 팔로잉 URL을 구함
	c <- result{url, name}              // 가져온 URL과 사용자 이름을 구조체 인스턴스로
                                            // 생성하여 채널 c에 보냄
}


// 실제 작업을 처리하는 worker 함수
func worker(done <-chan struct{}, urls chan string, c chan<- result) {
	for url := range urls { // urls 채널에서 URL을 가져옴
		select {
		case <-done: // 채널이 닫히면 worker 함수를 빠져나옴
			return
		default:
			crawl(url, urls, c) // URL 처리
		}
	}
}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs) // 모든 CPU를 사용하도록 설정

	urls := make(chan string)    // 작업을 요청할 채널
	done := make(chan struct{})  // 작업 고루틴에 정지 신호를 보낼 채널
	c := make(chan result)       // 결괏값을 저장할 채널

	var wg sync.WaitGroup
	const numWorkers = 10
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ { // 작업을 처리할 고루틴을 10개 생성
		go func() {
			worker(done, urls, c)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait() // 고루틴이 끝날 때까지 대기
		close(c)  // 대기가 끝나면 결괏값 채널을 닫음
	}()

	urls <- "https://github.com/pyrasis/following" // 최초 URL을 보냄

	count := 0
	for r := range c { // 결과 채널에 값이 들어올 때까지 대기한 뒤 값을 가져옴
		fmt.Println(r.name)

		count++
		if count > 100 {    // 100명만 출력한 뒤 
			close(done) // done을 닫아서 worker 고루틴을 종료
			break
		}
	}
}
