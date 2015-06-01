package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := make(map[string]interface{}) // 문자열을 키로하고 모든 자료형을 저장할 수 있는 맵 생성

	data["name"] = "maria"
	data["age"] = 10

	doc, _ := json.Marshal(data) // 맵을 JSON 문서로 변환

	fmt.Println(string(doc)) // {"age":10,"name":"maria"}: 문자열로 변환하여 출력
}
