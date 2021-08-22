package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func foo1() {
	time.Sleep(2 * time.Second)
	return
}

func handleRequest(timeout time.Duration) bool {
	ctx, cancle := context.WithCancel(context.Background())

	fooChan := make(chan struct{}, 1)
	go func() {
		foo1()
		fooChan <- struct{}{}
	}()

	//timeoutChan :=make(chan struct{},1)
	go func() {
		time.Sleep(timeout)
		cancle()
	}()

	select {
	case <-fooChan:
		fmt.Println("foo1 finish")
	case <-ctx.Done():
		fmt.Println("timeout")
	}

	return true
}

func main() {
	//handleRequest(1 * time.Second)
	//ch:=make(chan int)
	//close(ch)
	//close(ch)
	//close(ch)
	ticker := time.NewTimer(1 * time.Second)
	defer ticker.Stop()
	go func() {
		<-ticker.C
		fmt.Printf("go routine number %d\n", runtime.NumGoroutine())
	}()

	for {
		handleRequest(50 * time.Millisecond)
	}

}
