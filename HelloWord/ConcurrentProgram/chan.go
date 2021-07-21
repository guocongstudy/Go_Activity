package main

import "fmt"

/*创建管道*/

func main() {
	//定义一个管道,管道里面需要放类型，所以string不能掉
	var notice chan string = make(chan string)
	fmt.Printf("%T,%v\n", notice, notice)
	go func() {
		//往管道里面写数据
		notice <- "guocong"

		fmt.Println("start")
		//从管道里面读数据,并且进行输出
		fmt.Println(<-notice)

		fmt.Println("end")

	}()

}

//38:51
