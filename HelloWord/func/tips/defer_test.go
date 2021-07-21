package main

import (
	"fmt"
	"testing"
)

func TestDeferN(t *testing.T) {
	testDefer()
}

func testDefer() {
	fmt.Println("start")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")
}

func TestDefer11(t *testing.T) {
	x := 10
	//匿名函数
	defer func(m *int) {
		*m = 20
	}(&x) //（这个是传的实参，例子能直接看明白）
	x = 30
	fmt.Println(x)

	//匿名函数，一个例子看明白后面的括号是什么作用
	func(a int, b int) {
		fmt.Println("a=", a, "b=", b)
		return
	}(3, 5)
}
