package main

import (
	"fmt"
	"io"

	"os"
)

func main() {
	//定义文件路径
	path := "user.txt"
	//打开文件
	file, err := os.Open(path)
	fmt.Println(file)
	//判断文件
	if err != nil {
		fmt.Println(err)
	} else {
		//如果没有错误就定义一个字节切片
		//var bytes []byte = make([]byte, 20)
		//可以改成简短声明的方式
		bytes := make([]byte, 20)
		//用循环读取出文件
		for {
			n, err := file.Read(bytes)
			//判断错误，看err等不等于nil的变量
			if err != nil {
				//当err等于EOF的时候
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Println(string(bytes[:n]))
			}
		}
		file.Close()
	}
}
