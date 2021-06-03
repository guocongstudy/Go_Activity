package main

import "fmt"

type Address2 struct {
	Region string
	street string
	No     string
}

type User5 struct {
	ID   int
	Name string
	Addr Address2
}

type Employee struct {
	User5
	Salary float64
}

func main() {
	var me Employee
	fmt.Printf("%T,%#v", me, me)

	me02 := Employee{
		User5: User5{
			ID:   1,
			Name: "kk",
			Addr: Address2{"西安", "锦业路", "001"},
		},
		Salary: 1000,
	}
	fmt.Printf("%T,%#v\n", me, me02)
	fmt.Println(me02.Addr.street)
	me02.Name = "郭聪"
	fmt.Println(me02)

	//时间：1:04:37

}
