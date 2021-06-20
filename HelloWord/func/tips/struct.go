package main

import (
	"fmt"
	"time"
)

type User2 struct {
	ID       int
	Name     string
	Age      uint
	Birthday time.Time
}

func main() {
	var me User2

	u := User2{1, "guocong", 23, time.Now()}
	fmt.Println(u)
	fmt.Printf("%#v", me)
	fmt.Printf("%#v\n", me)

	//var me2 User2 = User2{2, "guocong", 33, time.Now()}
	//fmt.Printf("%#v\n", me2)

	var pointer *User2
	fmt.Printf("%T\n", pointer)
	fmt.Printf("%#v\n", pointer)

	var pointer2 *User2 = &User2{}
	fmt.Printf("%#v\n", pointer2)

}
