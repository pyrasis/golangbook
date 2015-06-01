package main

import "fmt"

func main() {
	solarSystem := make(map[string]float32) // 키는 string, 값은 float32인 맵 생성 및 공간 할당

	solarSystem["Mercury"] = 87.969 // 맵[키] = 값
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528

	for key, value := range solarSystem { // 반복문이 실행될 때마다 키와 값이 자동으로 변수에 들어감
		fmt.Println(key, value)
	}
}
