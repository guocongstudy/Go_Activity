package main

import (
	"flag"
	"fmt"
)

func main() {

	//绑定命令行参数与变量关系
	port := flag.Int("p", 22, "ssh port")
	host := flag.String("H", "127.0.0.1", "ssh host")
	verbor := flag.Bool("v", false, "detail log")
	help := flag.Bool("h", false, "help")

	flag.Usage = func() {
		fmt.Println("usage flag [-H 127.0.0.14][-p 22][-v]")
		flag.PrintDefaults()
	}

	//解析命令行参数
	flag.Parse()
	fmt.Printf("%T,%T,%T,%T\n", port, host, verbor, help)
	//解引用
	fmt.Printf("%T,%T,%T,%T", *port, *host, *verbor, *help)
}
