package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // 개인 키와 공개키 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey // 개인 키 변수 안에 공개 키가 들어있음

	s := `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	ciphertext, err := rsa.EncryptPKCS1v15( // 평문을 공개 키로 암호화
		rand.Reader,
		publicKey, // 공개키
		[]byte(s),
	)

	fmt.Printf("%x\n", ciphertext)

	plaintext, err := rsa.DecryptPKCS1v15( // 암호화된 데이터를 개인 키로 복호화
		rand.Reader,
		privateKey, // 개인키
		ciphertext,
	)

	fmt.Println(string(plaintext))
}
