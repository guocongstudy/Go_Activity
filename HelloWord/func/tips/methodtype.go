package main

import (
	"fmt"
)

type Pig struct {
	name string
}

func (pig Pig) Call() {
	fmt.Printf("%s:哼哼哼", pig.name)
}

func (pig Pig) SetName(name string) {
	pig.name = name
	//如果有返回值的话就应该这么写
	//return errors.New("输入错误").Error()
}

func main() {
	pig := Pig{"小猪"}
	m1 := pig.Call //卧槽，这里不能放括号，找半天找不出来错  pig拷贝
	fmt.Printf("%T\n", m1)
	pig.SetName("小黑")
	m1()

	ppig := &Pig{"豆豆"}
	m2 := ppig.Call //会自动解引用，拷贝。解引用以后就是指类型
	fmt.Printf("%T\n", m2)
	m2()
}

//00:17:47
