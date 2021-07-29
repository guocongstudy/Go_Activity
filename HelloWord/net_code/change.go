package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int, 1)
	num := 1
	channel <- num

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		a := <-channel
		fmt.Println(a)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(num)

}

//31讲，第一集
//《图解HTTP》，《HTTP权威指南》
