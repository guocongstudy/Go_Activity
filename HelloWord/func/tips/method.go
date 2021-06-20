package main

import "fmt"

type Dog struct {
	name string
}

func (dog Dog) Call() {
	fmt.Printf("%s:汪汪汪\n", dog.name)
}

//两个括号的意义：接收参数、函数参数
func (dog Dog) SetName(name string) {
	dog.name = name
}
func (dog *Dog) PsetName(name string) {
	dog.name = name
}
func main() {
	dog := Dog{"豆豆"}
	dog.Call()
	//看下面这个传法就无效
	dog.SetName("小白")
	dog.Call()

	//(&dog).PsetName("小黑")取引用
	(*&dog).PsetName("小蓝") //自动取引用
	dog.Call()

	pdog := &Dog{"小豆"}
	(*pdog).Call()
	//(*pdog).Call() 解引用
	pdog.PsetName("小绿")
	(*pdog).Call()
}
