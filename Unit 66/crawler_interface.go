package main

import (
	"fmt"
	"golang.org/x/net/html" // HTML 파싱 해키지
	"log"
	"net/http"
	"runtime"
	"sync"
)

// 팔로잉 결괏값을 저장할 구조체
type FollowingResult struct {
	url  string
	name string
}

// 별표 저장소 결괏값을 저장할 구조체
type StarsResult struct {
	repo string
}

// 중복 URL을 처리할 구조체
type FetchedUrl struct {
	m map[string]error
	sync.Mutex
}

//  중복 저장소 이름을 처리할 구조체
type FetchedRepo struct {
	m map[string]struct{}
	sync.Mutex
}

// Crawler 인터페이스 정의
type Crawler interface {
	Crawl()
}

// 팔로잉 수집 구조체 정의
type GitHubFollowing struct {
	fetchedUrl *FetchedUrl          // 중복 URL을 처리할 멤버
	p          *Pipeline            // 파이프라인에 작업 요청을 보낼 멤버
	stars      *GitHubStars         // 별표 페이지도 처리해야 하므로 stars 멤버도 포함
	result     chan FollowingResult // 결괏값을 보낼 멤버
	url        string               // URL 처리 작업을 요청하기 위한 멤버
}

// 별표 저장소 수집 구조체 정의
type GitHubStars struct {
	fetchedUrl  *FetchedUrl      // 중복 URL을 처리할 멤버
	fetchedRepo *FetchedRepo     // 중복 저장소 이름을 처리할 멤버
	p           *Pipeline        // 파이프라인에 작업 요청을 보낼 멤버
	result      chan StarsResult // 결괏값을 보낼 멤버
	url         string           // URL 처리 작업을 요청하기 위한 멤버
}

func fetch(url string) (*html.Node, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return doc, nil
}

func (g *GitHubFollowing) Request(url string) {
	g.p.request <- &GitHubFollowing{     // 구조체를 생성하여 파이프라인의 request 채널에 보냄
		fetchedUrl: g.fetchedUrl,    // 현재 인스턴스에서 포인터를 가져와서 다시 사용
		p:          g.p,             // 현재 인스턴스에서 포인터를 가져와서 다시 사용
		result:     g.result,        // 현재 인스턴스에서 포인터를 가져와서 다시 사용
		stars:      g.stars,         // 현재 인스턴스에서 포인터를 가져와서 다시 사용
		url:        url,             // url만 새로운 값
	}
}

// 팔로잉 페이지에서 사용자 이름, 팔로잉 URL, 별표 URL을 구함
func (g *GitHubFollowing) Parse(doc *html.Node) <-chan string {
	name := make(chan string)

	go func() {
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "img" {
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

					if a.Key == "class" && a.Val == "avatar" {
						// 부모 요소의 첫 번째 속성(href)
						user := n.Parent.Attr[0].Val

						// 팔로잉 URL을 파이프라인에 보냄
						g.Request("https://github.com" + user + "/following")
						// 별표 URL을 파이프라인에 보냄
						g.stars.Request("https://github.com/stars" + user)
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

	return name
}

func (g *GitHubFollowing) Crawl() {
	g.fetchedUrl.Lock()                     // 맵은 뮤텍스로 보호
	if _, ok := g.fetchedUrl.m[g.url]; ok { // URL 중복 처리 여부를 검사
		g.fetchedUrl.Unlock()
		return
	}
	g.fetchedUrl.Unlock()

	doc, err := fetch(g.url)     // URL에서 파싱된 HTML 데이터를 가져옴
	if err != nil {              // URL을 가져오지 못했을 때
		go func(u string) {  // 교착 상태가 되지 않도록 고루틴을 생성
			g.Request(u) // 파이프라인에 URL을 보냄
		}(g.url)
		return
	}

	g.fetchedUrl.Lock()
	g.fetchedUrl.m[g.url] = err // 가져온 URL은 맵에 URL과 에러 값 저장
	g.fetchedUrl.Unlock()

	name := <-g.Parse(doc)                   // 사용자 이름을 구함
	g.result <- FollowingResult{g.url, name} // 사용자 이름과 URL을 팔로잉 결과 채널에 보냄
}

func (g *GitHubStars) Request(url string) {
	g.p.request <- &GitHubStars{        // 구조체를 생성하여 파이프라인의 request 채널에 보냄
		fetchedUrl:  g.fetchedUrl,  // 현재 인스턴스에서 포인터를 가져와서 다시 사용
		fetchedRepo: g.fetchedRepo, // 현재 인스턴스에서 포인터를 가져와서 다시 사용
		p:           g.p,           // 현재 인스턴스에서 포인터를 가져와서 다시 사용
		result:      g.result,      // 현재 인스턴스에서 포인터를 가져와서 다시 사용
		url:         url,           // url만 새로운 값
	}
}

// 별표 페이지에서 저장소 이름을 구함
func (g *GitHubStars) Parse(doc *html.Node) <-chan string {
	repo := make(chan string)

	go func() {
		defer close(repo) // 고루틴이 끝나기 직전에  채널을 닫음

		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "span" { // span 태그
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "prefix" { // class가 prefix인 요소
						// 저장소 이름을 구한 뒤 채널에 보냄
						repo <- n.Parent.Attr[0].Val 
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

	return repo
}

func (g *GitHubStars) Crawl() {
	g.fetchedUrl.Lock()                     // 맵은 뮤텍스로 보호
	if _, ok := g.fetchedUrl.m[g.url]; ok { // URL 중복 처리 여부를 검사
		g.fetchedUrl.Unlock()
		return
	}
	g.fetchedUrl.Unlock()

	doc, err := fetch(g.url)     // URL에서 파싱된 HTML 데이터를 가져옴
	if err != nil {              // URL을 가져오지 못했을 때
		go func(u string) {  // 교착 상태가 되지 않도록 고루틴을 생성
			g.Request(u) // 파이프라인에 URL을 보냄
		}(g.url)
		return
	}

	g.fetchedUrl.Lock()
	g.fetchedUrl.m[g.url] = err // 가져온 URL은 맵에 URL과 에러 값 저장
	g.fetchedUrl.Unlock()

	repositories := g.Parse(doc)
	for r := range repositories {                   // 채널에서 range로 저장소 이름을 계속 꺼냄
		g.fetchedRepo.Lock()                    // 맵은 뮤텍스로 보호
		if _, ok := g.fetchedRepo.m[r]; !ok {   // 저장소 중복 저장 여부 검사
			g.result <- StarsResult{r}      // 저장소 이름을 별표 저장소 결과 채널에 보냄
			g.fetchedRepo.m[r] = struct{}{} // 저장한 저장소 이름은 맵에 저장
		}
		g.fetchedRepo.Unlock()
	}
}

// Crawler 인터페이스를 처리할 파이프라인 구조체 정의
type Pipeline struct {
	request chan Crawler // Crawler 타입으로 채널 선언
	done    chan struct{}
	wg      *sync.WaitGroup
}

func NewPipeline() *Pipeline {
	return &Pipeline{ // 새 파이프라인 생성
		request: make(chan Crawler),
		done:    make(chan struct{}),
		wg:      new(sync.WaitGroup),
	}
}

// 실제 작업을 처리하는 worker 함수
func (p *Pipeline) Worker() {
	for r := range p.request { // request 채널에서 Crawler 인터페이스를 가져옴
		select {
		case <-p.done: // 채널이 닫히면 worker 함수를 빠져나옴
			return
		default:
			r.Crawl() // Crawler 인터페이스의 Crawl 함수 실행
		}
	}
}

func (p *Pipeline) Run() {
	const numWorkers = 10
	p.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ { // 작업을 처리할 고루틴 10개 생성
		go func() {
			p.Worker()
			p.wg.Done()
		}()
	}

	go func() {
		p.wg.Wait() // 작업 고루틴이 끝날 때까지 대기
	}()
}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs) // 모든 CPU를 사용하도록 설정

	p := NewPipeline() // 파이프라인 인스턴스 생성
	p.Run()            // worker 함수를 고루틴으로 생성

	stars := &GitHubStars{                                           // 별표 구조체 인스턴스 생성
		fetchedUrl:  &FetchedUrl{m: make(map[string]error)},     // 중복 URL 처리 맵
		fetchedRepo: &FetchedRepo{m: make(map[string]struct{})}, // 중복 저장소 
                                                                         // 이름 처리 맵
		p:           p,                                          // 파이프라인 인스턴스
		result:      make(chan StarsResult),                     // 별표 저장소 
                                                                         // 결괏값 채널
	}
	following := &GitHubFollowing{                              // 팔로잉 구조체 인스턴스 생성
		fetchedUrl: &FetchedUrl{m: make(map[string]error)}, // 중복 URL 처리 맵
		p:          p,                                      // 파이프라인 인스턴스
		result:     make(chan FollowingResult),             // 팔로잉 결괏값 채널
		stars:      stars,                                  // GitHubStars 구조체 인스턴스
		url:        "https://github.com/pyrasis/following", // 최초 시작 URL
	}
	p.request <- following // 파이프라인에 구조체 인스턴스를 보내서 HTML 처리 작업 시작

	count := 0
LOOP:
	for { // 무한 루프 사용
		select { // 채널에 결괏값이 들어왔을 때마다 값을 출력
		case f := <-following.result:
			fmt.Println(f.name)
		case s := <-stars.result:
			fmt.Println(s.repo)
			if count == 1000 {     // 저장소를 1,000개만 출력
				close(p.done)  // done 채널을 닫음
				break LOOP     // 무한 루프를 빠져나옴
			}
			count++
		}
	}
}
