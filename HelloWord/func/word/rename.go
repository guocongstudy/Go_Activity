package main

import "os"

func main() {
	//给文件更换名字，多运行几次，第一次看可能没换好。
	os.Rename("user.log", "user2.log")
	//os.Remove("user2.log")
}
