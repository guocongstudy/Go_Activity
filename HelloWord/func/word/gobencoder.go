package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
}

func main() {
	users := map[int]User{
		1: User{1, "烟灰", time.Now(), "福建", "1213323"},
		2: User{2, "郭聪", time.Now(), "汉中", "3243423"},
		3: User{3, "小狗", time.Now(), "渭南", "2323232"},
	}

	file, err := os.Create("user.gob")
	if err == nil {
		defer file.Close()

		encoder := gob.NewEncoder(file)
		encoder.Encode(&users)

		fmt.Printf("%#v", users)

	}
}
