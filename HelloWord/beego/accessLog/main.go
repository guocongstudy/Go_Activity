package main

import (
	"fmt"
	"os"
)

func main() {
	logFile, _ := os.Create("test.log")
	fmt.Fprintln(logFile, "日志1")
	fmt.Println("日志2")
	fmt.Println("日志3")

	//log包
}
