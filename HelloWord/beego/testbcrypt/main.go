package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "123@456"
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	fmt.Println(hashed, err)

}
