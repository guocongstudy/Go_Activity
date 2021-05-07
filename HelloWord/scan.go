package main

import "fmt"

func main(){
	var name string
	fmt.Println("请输入名称：")
	fmt.Scan(&name)
	fmt.Println("你输入的是:",name)
	var age int
	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age)
	fmt.Println("年龄：",age)
}
/*
笔记：
\\：反斜线
\':单引号
\":双引号
\a:响铃
\b:退格
\f:换页
\n:换行
\r:回车
\t:制表符
\v:垂直制表符
*/