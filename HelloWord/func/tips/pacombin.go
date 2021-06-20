package main

import "fmt"

type Address9 struct {
	Region string
	Street string
	No     string
}

type User9 struct {
	ID   int
	Name string
	Addr Address9
}

type Employee9 struct {
	*User9
	Salary float64
	Name   string
}

func main() {
	var me Employee
	fmt.Printf("%#v\n", me)
	user := User9{
		ID:   1,
		Name: "郭聪",
		Addr: Address9{"西安市", "太白路", "玫瑰公馆"},
	}
	employee := Employee9{
		User9: &User9{ //这一步特别注意，我自己写的时候User：没有写 红一片
			ID:   1,
			Name: "guoocong",
			Addr: Address9{"北京市", "海淀区", "学院南路"},
		},
		Salary: 12000,
		Name:   "小郭",
	}
	fmt.Println(user)
	fmt.Println(employee)
}
