package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddFunc(t *testing.T) {
	x, y := 2, 4
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)

	//匿名函数
	add := func(x, y int) int {
		return x + y
	}
	add(1, 3)
}

//交换函数
func swap(x, y *int) {
	*x, *y = *y, *x
	return
}

//闭包

/*go实现了一级函数，函数是一个值，可以将函数赋值给变量。使得整个变量也成为函数，因此我们也可以定义函数
的类型可以将func作为一种type，以后可以直接使用这个type来定义函数。
*/
func Test11(t *testing.T) {
	var my MyType
	my = func(x, y int) int {
		return x + y
	}
	fmt.Println(reflect.TypeOf(my))
	my(1, 3)

	addFunc := func(x, y int) int {
		return x + y
	}
	fmt.Println(reflect.TypeOf(addFunc))
	fmt.Println(addFunc(10, 30))

	my1 := func(x, y int) int {
		return x + y
	}
	T11(my1)
}

//定义一个函数类型
type MyType func(x, y int) int

func T11(fn MyType) int {
	return fn(1, 2) * 2
}

//递归
func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}
