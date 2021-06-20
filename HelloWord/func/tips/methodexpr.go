package main

import "fmt"

//方法表达式
type Kiss struct {
	name string
}

func (kiss Kiss) Call() {
	fmt.Println("%s:嘶嘶嘶\n", kiss.name)
}

func (kiss *Kiss) SetName(name string) {
	kiss.name = name
}
func main() {
	//call定义的是值接受者
	m1 := Kiss.Call
	//SetName使用指针定义的接受者，所以有（*kiss）
	m2 := (*Kiss).SetName
	fmt.Println(m1, m2) //这样打印出来的是m1 m2的十六进制地址
	fmt.Printf("%T", m1)
	fmt.Printf("%T", m2)
	kiss := Kiss{"小盖"}
	m1(kiss)

	//585130764  7:30
}
