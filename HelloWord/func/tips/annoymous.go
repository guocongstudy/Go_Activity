package main

import "fmt"

func main() {
	//匿名结构体,直接用结构体声明一个变量
	var me struct {
		ID   int
		Name string
	}
	me.Name = "guocong"
	fmt.Printf("%T\n", me)
	fmt.Printf("%#v\n", me)
	fmt.Println(me.ID)
	//上下都是一个概念，两种不同的写法

	me2 := struct {
		ID   int
		Name string
	}{1, "kkk"}
	fmt.Printf("%#v\n", me2)
}
