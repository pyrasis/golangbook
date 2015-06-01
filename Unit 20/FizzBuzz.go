package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {                         // case에 조건식을 지정했으므로 판단할 변수는 생략
		case i%3 == 0 && i%5 == 0:       // 3의 배수이면서 5의 배수일 때
			fmt.Println("FizzBuzz")  // FizzBuzz 출력
		case i%3 == 0:                   // 3의 배수일 때
			fmt.Println("Fizz")      // Fizz 출력
		case i%5 == 0:                   // 5의 배수일 때
			fmt.Println("Buzz")      // Buzz 출력
		default:                         // 아무 조건에도 해당하지 않을 때
			fmt.Println(i)           // 숫자 출력
		}
	}
}
