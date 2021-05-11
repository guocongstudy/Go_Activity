package main

import "fmt"

func main(){
	var bytes []byte=[]byte{'A','B','C'}
	fmt.Printf("%T,%#v\n\n",bytes,bytes)

	s:=string(bytes)
	fmt.Printf("%T %v\n",s,s)

	bs :=[]byte(s)
	fmt.Printf("%T %#v\n",bs,bs)
}
