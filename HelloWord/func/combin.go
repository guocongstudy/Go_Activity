package main

import "fmt"

type Address struct {
	Region string
	Street string
	No     string
}

type User4 struct {
	ID   int
	Name string
	Addr Address
}

func main() {
	//me01是空结构体
	var me01 User4
	//User4{1,"郭聪",{"汉中","汉台区","原上村"}}
	addr := Address{"西安市", "太白路", "中段22号"}

	me02 := User4{
		ID:   1,
		Name: "kk",
		Addr: addr,
	}
	fmt.Println(me02)
	fmt.Println(me01)

	me03 := User4{
		ID:   2,
		Name: "guocong",
		Addr: Address{
			Region: "北京市",
			Street: "海淀区",
			No:     "002",
		},
	}
	fmt.Println(me03)
	fmt.Printf("%#v\n", me03)
	fmt.Println(me03.Addr.Region)
}
