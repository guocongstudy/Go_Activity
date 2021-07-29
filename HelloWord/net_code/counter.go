package main

import "time"

func main() {
	//计数例程的数量
	var stopCounter int = 0
	const counter = 3
	//3个例程
	for i := 0; i < counter; i++ {
		go func() {
			stopCounter++
		}()
	}
	//如何等待stopCounter是3
	//让不等于3的时候程序不能退出
	for stopCounter != counter {
		time.Sleep(3 * time.Second)
	}
}
