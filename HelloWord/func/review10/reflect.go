package main

import (
	"fmt"
	"reflect"
)

//最简单的反射

func main() {
	var i int = 1
	//反射形式1：
	fmt.Printf("%T\n", i)
	//反射形式2：
	var typ = reflect.TypeOf(i)
	fmt.Println(typ)
}
