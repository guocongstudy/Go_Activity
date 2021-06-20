package main

import "fmt"

type Class struct {
	id   int
	name string
}

//getName不管用值接收还是指针接收都可以，但是setName一定要用指针接收
func (class Class) GetID() int {
	return class.id
}

func (class Class) GetName() string {
	return class.name
}

func (class *Class) SetID(id int) int {
	class.id = id
	return class.id
}

func (class *Class) SetName(name string) string {
	class.name = name
	return class.name
}

type worker struct {
	Class
	salary float64
	name   string
}

func main() {
	var me worker = worker{ //前后两个work要大小单词都写一样才行
		Class: Class{
			id:   1,
			name: "郭聪",
		},
		salary: 10000,
		name:   "小盖",
	}
	var c Class = Class{ //前后两个class要大小单词都写一样才行
		id:   1,
		name: "guocong",
	}
	fmt.Println(c)
	fmt.Println(me)
	fmt.Println(me.id)
	fmt.Println(me.name)
	fmt.Println(me.Class.GetName())
	fmt.Println(me.Class.GetID())
	me.SetName("silence")
	fmt.Println(me.GetName())

}
