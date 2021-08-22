package main

import (
	"fmt"
	"os"
	"time"
)

func CountDown(n int, countChan chan int, finishChan chan struct{}) {
	if n <= 0 {
		return
	}
	//每隔一秒钟减一
	ticker := time.NewTimer(1 * time.Second)
	for {
		countChan <- n
		//等待一秒钟
		<-ticker.C
		n--
		//退出循环的条件
		if n <= 0 {
			//关掉ticker
			ticker.Stop()
			//表示已经数完了
			finishChan <- struct{}{}
			break
		}
	}
}

//监听键盘有没有输入
func Abort(abortChan chan struct{}) {
	bs := make([]byte, 1)
	//从标准输入中读一行出来
	os.Stdin.Read(bs)
	abortChan <- struct{}{}
}

func main() {
	//创建一下管道
	countChan := make(chan int)
	finishChan := make(chan struct{})
	abortChan := make(chan struct{})

	go CountDown(10, countChan, finishChan)
	go Abort(abortChan)
LOOP:
	for {
		select {
		case n := <-countChan:
			fmt.Println(n)
		case <-finishChan:
			fmt.Println("biu biu biu")
			break LOOP
		case <-abortChan:
			fmt.Println("abort")
			break LOOP
		}
	}

}
