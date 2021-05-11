package main

import (
	"fmt"
	"strconv"
)

func main(){
	//字符串=>其他类型
	//=>bool

	if v,err :=strconv.ParseBool("true");err ==nil{
		fmt.Println(v)
	}else{
		fmt.Println(err)
	}

	//int
	if v,err :=strconv.Atoi("1023");err ==nil{
		fmt.Println(v)
	}else{
		fmt.Println(err)
	}
	if v,err :=strconv.ParseInt("64",16,64);err ==nil{
		fmt.Println(v)
	}else{
		fmt.Println(err)
	}
	//float
	if v,err :=strconv.ParseFloat("1.1",64);err ==nil{
		fmt.Println(v)
	}
}
