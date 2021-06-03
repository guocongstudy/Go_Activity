package main

import (
	"fmt"
	"time"
)

type Student struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Remark   string
}

func main() {
	me := Student{
		Name: "guocong",
		ID:   1,
		Addr: "西安",
		Tel:  "2132323",
	}
	fmt.Println(me.ID, me.Name, me.Tel)

	me.Tel = "12313200"
	fmt.Printf("%#v\n", me)
	me2 := &Student{
		ID:   2,
		Name: "woie",
	}
	fmt.Printf("%T\n", me2)
	(*me2).Tel = "1323344"   //这个写法也很牛逼，解引用
	fmt.Printf("%#v\n", me2) //%#v会给出实例的完整输出，包括他的字段（在程序自动生成go代码也很有用）
	me2.Addr = "北京"
	fmt.Printf("%#v\n", me2)

}
