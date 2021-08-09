package main

import "log"

//go里面也提供了log包
func main() {
	//log包
	//log的记录，默认记录在控制台上
	//在log里面，printf 和printfln其实没有什么区别
	//log会自动进行空格换行
	log.Panic("panic")
	log.Printf("print")
	//前缀,SetPrefix设置标准日志记录器的输出前缀。
	log.SetPrefix("testlog")
	log.Flags()
	log.SetFlags(log.Flags() | log.Lshortfile)
	//将日志进行输出
	//log.Output()

}

//beego的介绍与使用

//3：19：32
