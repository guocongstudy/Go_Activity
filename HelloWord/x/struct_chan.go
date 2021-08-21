package main

import "fmt"

var ch1 = make(chan struct{})

func subGroutine(ch chan struct{}){
	//do something
	fmt.Println("subGroutine finish")
	ch <- struct{}{}
	close(ch)
}
func main(){
	go subGroutine(ch1)
	<-ch1
}