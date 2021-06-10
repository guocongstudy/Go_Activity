package main

import (
	"fmt"
	"os"
)

func main() {
	path := "user2.txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	} else {
		//写文件
		file.Write([]byte("abc123!@#"))
		file.WriteString("guocong")
		//关闭文件,关闭文件指令后面不能再加写文件，关了以后是加不进去的
		file.Close()

	}
}
