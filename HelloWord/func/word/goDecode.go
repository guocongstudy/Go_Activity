package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type User2 struct {
	ID       int
	Name     string
	Birthday time.Time
	Tel      string
	Addr     string
}

func main() {
	var std User2

	file, err := os.Open("user.gob")
	if err == nil {
		defer file.Close()

		decoder := gob.NewDecoder(file)
		//解码
		decoder.Decode(&std)
		fmt.Printf("%#v", std)
	}
}
