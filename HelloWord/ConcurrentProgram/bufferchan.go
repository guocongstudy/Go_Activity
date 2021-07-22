package main

import (
	"fmt"
	"time"
)

func main() {
	//带缓冲的管道，提前往里面写两个值，就不会阻塞
	channel := make(chan string, 2)
	channel <- "x"
	channel <- "y"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println("--------------------------------")

	var channel1 chan int = make(chan int, 5)
	//只读管道(在前面)
	var rchannel <-chan int = channel1
	//只写管道(在后面)
	var wchannel chan<- int = channel1

	go func(channel1 <-chan int) {
		fmt.Println(<-channel1)
	}(channel1)

	go func(channel1 chan<- int) {
		channel1 <- 0
	}(channel1)

	wchannel <- 1
	fmt.Println(<-rchannel)
	time.Sleep(3 * time.Second)
}
