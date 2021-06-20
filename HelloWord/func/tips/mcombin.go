package main

import (
	"fmt"
)

type Addresss struct {
	Region string
	Street string
	No     string
}

func (addr Addresss) String() string {
	return fmt.Sprintf("%s-%s-%s", addr.Region, addr.Street, addr.No)
}

type Users struct {
	ID   int
	Name string
	Addr Addresss
}

func (users Users) String() string {
	return fmt.Sprintf("[%d]%s:%s", users.ID, users.Name, users.Addr)
}

func main() {
	var u Users = Users{
		ID:   1,
		Name: "kk",
		Addr: Addresss{"西安市", "锦业路", "001"},
	}
	fmt.Println(u)
	fmt.Println(u.Addr)
}
