package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := "Hello, Key 12345" // 16바이트
	s := "Hello, world! 12"   // 16바이트

	block, err := aes.NewCipher([]byte(key)) // AES 대칭키 암호화 블록 생성
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s)) // 평문을 AES 알고리즘으로 암호화
	fmt.Printf("%x\n", ciphertext)

	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext) // AES 알고리즘으로 암호화된 데이터를 평문으로 복호화
	fmt.Println(string(plaintext))
}
