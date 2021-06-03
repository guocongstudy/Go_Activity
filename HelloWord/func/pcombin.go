package main

import "fmt"

type Address8 struct {
	Region string
	Street string
	No     string
}

type User8 struct {
	ID   int
	Name string
	Addr *Address8
}

func main() {
	var me01 User8
	fmt.Println("%#v", me01)
	me02 := User8{
		ID:   1,
		Name: "kk",
		Addr: &Address8{"西安市", "鄠邑路", "西安石油大学"},
	}
	fmt.Printf("%#v\n", me02)
	fmt.Printf("%T\n", me02.Addr)
	fmt.Println(me02.Addr.Region)
	me02.Addr.Region = "上海市"
	fmt.Printf("%#v\n", me02)
	fmt.Println(me02)
	fmt.Println(me02.Addr.Region)

}
