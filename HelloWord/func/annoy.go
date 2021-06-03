package main

import (
	"fmt"
)

type Address3 struct {
	Region string
	Street string
	No     string
}

type User6 struct {
	ID   int
	Name string
	Addr *Address3
}

func main() {
	var me2 User6
	fmt.Printf("%#v\n", me2)
	me3 := User6{
		ID:   1,
		Name: "kk",
		Addr: &Address3{"西安市", "太白路", "dw"},
	}
	fmt.Printf("%#v\n", me3)
	fmt.Println(me3)

}
