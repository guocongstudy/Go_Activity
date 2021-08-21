package main

import (
	"fmt"
	"time"
)

func grandson(){
	fmt.Println("grandson begin")
	time.Sleep(5*time.Second)
	fmt.Printf("grandson finish")
}


func child(){
	fmt.Println("child begin")
	go grandson() //子协程调用孙携程
	time.Sleep(1*time.Second)
	fmt.Printf("child finish")
}

func main(){
	go child()
	time.Sleep(10*time.Second)
}