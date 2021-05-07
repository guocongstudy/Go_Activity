package main

import "fmt"

func main()  {
	fmt.Println("老婆的想法：")
	fmt.Println("十个包子")
	var yes string
	fmt.Println("有卖西瓜的吗？(y/Y)")
	fmt.Scan(&yes) //很容易掉指针符号
	switch yes{
	case "y":
		fmt.Println("买一个西瓜")
	case "Y":
		fmt.Println("买一个西瓜")
	default://默认是什么
		fmt.Println("输入错误")
	}
	var result = 0

	//循环累加到100
	for i :=1 ; i<=100 ; i++{
		result += i
	}
	fmt.Println(result)

	//方法二
	result = 0
	i :=1
	for i <=100 {
		result +=i
		i++
	}
	fmt.Println(result)

	//方法三
	desc := "我爱中国"
	for j,ch:=range desc{
		fmt.Printf("%d %T %q\n",j,ch,ch)
	}
}
