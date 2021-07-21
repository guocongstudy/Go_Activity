package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
	sex  int    `json:"sex,int,omitempty"` //omitempty零值省略
	tel  string `json:"-"`
	Addr string
}

func main() {
	User := User{1, "kk", 1, "152000", "西安"}
	bytes, _ := json.MarshalIndent(User, "", "\t")
	fmt.Println(string(bytes))
}
