package main

import "fmt"

type Duck struct { // 오리(Duck) 구조체 정의
}

func (d Duck) quack() {     // 오리의 quack 메서드 정의
	fmt.Println("꽥~!") // 오리 울음 소리
}

func (d Duck) feathers() { // 오리의 feathers 메서드 정의
	fmt.Println("오리는 흰색과 회색 털을 가지고 있습니다.")
}

type Person struct { // 사람(Person) 구조체 정의
}

func (p Person) quack() {                           // 사람의 quack 메서드 정의
	fmt.Println("사람은 오리를 흉내냅니다. 꽥~!") // 사람이 오리 소리를 흉내냄
}

func (p Person) feathers() { // 사람의 feathers 메서드 정의
	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")
}

type Quacker interface { // quack, feathers 메서드를 가지는 Quacker 인터페이스 정의
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()    // Quacker 인터페이스로 quack 메서드 호출
	q.feathers() // Quacker 인터페이스로 feathers 메서드 호출
}

func main() {
	var donald Duck // 오리 인스턴스 생성
	var john Person // 사람 인스턴스 생성

	inTheForest(donald) // 인터페이스를 통하여 오리의 quack, feather 메서드 호출
	inTheForest(john)   // 인터페이스를 통하여 사람의 quack, feather 메서드 호출
}
