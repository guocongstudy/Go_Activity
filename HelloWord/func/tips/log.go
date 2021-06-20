package main

import "log"

//日志管理
func main() {
	log.Printf("我是printf日志：%s", "x")
	//用于日志分类，给日志加前缀
	log.SetPrefix("prefix:")
	log.SetFlags(log.Flags() | log.Lshortfile)
	log.Printf("我是printf日志：%s", "x")

	//log.Panicf("我是panic日志：%s"，"y")
	log.Fatalf("我是Fatalf日志：%s", "z")
	log.Panicf("我是panic日志：%s", "y")
}

/* 作业：
1.整理知识点
2.用户管理 做完

3.创建一个module
 可运行
 用户管理放在子包中（库程序）
4.密码铭文，命令行输入铭文
 密码输入：gopass
 代码中的密码 使用MD5，并在验证时计算输入md5值并进行比较


*/
