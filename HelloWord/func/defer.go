package main

import "fmt"

func main(){
	defer func(){
	fmt.Println("defer01")
	}()
	defer func(){ //当函数退出的时候执行
		fmt.Println("defer02")
	}()
	fmt.Println("main over")
}
