package main

import "fmt"

var ch2 =make(chan bool,10)

func consumer(ch2 chan bool){
	for {
		v:=<-ch2
		fmt.Printf("take %t\n",v)

	}
}

func producer(ch2 chan bool){
	for i:=0;i<10;i++{
		ch2 <-true
	}
}


func main(){
	producer(ch2)
	close(ch2)
	for els :=range ch {
		fmt.Printf("read %t\n",els)
	}
}