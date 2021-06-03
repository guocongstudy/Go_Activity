package main

import "fmt"

type User7 struct {
	ID   int
	Name string
	Addr Address4
}
type Address4 struct {
	ID   int
	Name string
	No   string
}
type Employee2 struct {
	//指针类型的匿名嵌入
	*User7
	Salary float64
	Name   string
}

func main() {
	var me Employee2
	fmt.Printf("%#v\n", me)

	me5 := Employee2{
		User7: &User7{
			ID:   1,
			Name: "kk",
			Addr: Address4{1, "锦业路", "001"},
		},
		Salary: 1200,
		Name:   "小K",
	}
	fmt.Printf("%#v\n", me5)
	fmt.Println(me5.Name)
	fmt.Println(me5.Addr)

}
