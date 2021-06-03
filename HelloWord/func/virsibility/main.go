package main

import (
	"fmt"
	"user/users"
)

//5.结构体可加性
//结构体的可见性
func main() {
	var u users.User

	fmt.Println(u)
	fmt.Println(u.Name)
	fmt.Println(u.ID)
}

/*
匿名嵌套
结构体首字母大写包外可见，否则仅包内可见。
结构体名字：
StructName
structName
属性：
AttrName
attrName
结构体名 结构体属性 => 结构体
S A => s A=>结构名小写，包外不能创建结构体对象
S a => s a=>结构体名小写，包外不能创建结构体对象

S A => S A=>结构体名大小，包外可以创建结构体对象，属性也可以访问
S a => S a=>结构体名大小，包外可以创建结构体对象，属性也不可以访问

s A => S A=>结构体名大小，包外可以创建结构体对象，属性也可以访问
s a => S a=>结构体名大小，包外可以创建结构体对象，属性也不可以访问

s A => s A=>结构名小写，包外不能创建结构体对象
s a => s a=>结构体名小写，包外不能创建结构体对象
*/
