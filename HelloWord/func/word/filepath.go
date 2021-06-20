package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//绝对路径
	fmt.Println(filepath.Abs("."))
	fmt.Println(os.Args)
	//path,_:=fmt.Println(filepath.Abs(os.Args[0]))

	fmt.Println(filepath.Base("c:/test/a.txt"))
	fmt.Println(filepath.Base("c:/test/xxxx/"))
	//fmt.Println(filepath.Dir("path"))

	fmt.Println(filepath.Ext("c:/test/a.txt"))
	fmt.Println(filepath.Ext("c:/test/xxx/a"))
	//fmt.Println(filepath.Ext("path"))

}
