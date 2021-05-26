package main

import (
	"flag"
	"fmt"
)

func main() {
	var host string
	var port int
	var verbor bool
	var help bool
	//绑定命令行参数与变量关系
	flag.IntVar(&port, "p", 22, "ssh port")
	flag.StringVar(&host, "H", "127.0.0.1", "ssh host")
	flag.BoolVar(&verbor, "v", false, "detail log")
	flag.BoolVar(&help, "h", false, "help")
	flag.Usage = func() {
		fmt.Println("usage flag [-H 127.0.0.14][-p 22][-v]")
		flag.PrintDefaults()
	}

	//解析命令行参数
	flag.Parse()
	if help {
		flag.Usage()
	} else {
		fmt.Println(host, port, verbor)
	}

	/*问题：
	  flag.IntVar这个函数的意思？
	*/

}
