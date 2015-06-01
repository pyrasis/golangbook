package main

import (
	"fmt"
	"regexp"
)

func main() {
	var validEmail, _ = regexp.Compile(
		"^[_a-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$",
	)

	fmt.Println(validEmail.MatchString("hello@mail.example.com"))    // true
	fmt.Println(validEmail.MatchString("hello+kr@example.com"))      // true
	fmt.Println(validEmail.MatchString("hello-world@example.co.kr")) // true
	fmt.Println(validEmail.MatchString("hello_1@example.info"))      // true
	fmt.Println(validEmail.MatchString("gildong.hong@e-xample.com")) // true

	fmt.Println(validEmail.MatchString("@example.com"))            // false
	fmt.Println(validEmail.MatchString("hello@example"))           // fasle
	fmt.Println(validEmail.MatchString("hello@example.cooooooom")) // false
}
