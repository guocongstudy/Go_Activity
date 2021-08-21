package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 10)
//这个是输出
func cons(ch chan int) {
	for i := 0; i < 5; i++ {
		v := <-ch
		fmt.Printf("take %d\n", v)
	}
}
//这个是存入
func prod (ch chan int){
	for i:=0;i<11;i++{
		ch <-i
		fmt.Printf("SEND %d\n",i)
	}
}

func main(){
	//先存入再输出
	go prod(ch)
	go cons(ch)
	time.Sleep(2*time.Second)
}