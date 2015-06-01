package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	doc := `
	{
		"name": "maria",
		"age": 10
	}
	`

	var data map[string]interface{} // JSON 문서의 데이터를 저장할 공간을 맵으로 선언

	json.Unmarshal([]byte(doc), &data) // doc를 바이트 슬라이스로 변환하여 넣고, 
                                           // data의 포인터를 넣어줌

	fmt.Println(data["name"], data["age"]) // maria 10: 맵에 키를 지정하여 값을 가져옴
}
