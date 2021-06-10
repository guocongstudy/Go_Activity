package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("user.txt")

	bytes := make([]byte, 100)
	n, _ := file.Read(bytes)
	n, err := file.Read(bytes)

	fmt.Println(n, err)
	file.Close()
}
