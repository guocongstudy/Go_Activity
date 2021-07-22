package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//一个操作3s为执行完成就任务超时
	result := make(chan int)
	timeout := make(chan int)
	//任务例程
	go func() {
		time.Sleep(time.Duration(rand.Int()%10) * time.Second)
		result <- 0
	}()

	go func() {
		time.Sleep(3 * time.Second)
		timeout <- 0
	}()

	select {
	case <-result:
		fmt.Println("执行完成")
	case <-timeout:
		fmt.Println("执行超时")
	}

}

//00:38:03
