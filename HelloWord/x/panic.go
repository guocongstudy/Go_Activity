package main

import "fmt"

func goo(x int) int {
	fmt.Printf("x=%d\n", x)
	return x
}

func foo(a, b int) int {
	c := a*3 + 8
	defer fmt.Printf("1st defer")
	d := c + 5
	defer fmt.Println("2nd defer")
	e := d / b
	defer fmt.Println("3rd defer")
	return goo(e)
}

func main() {
	foo(3, 4)
	fmt.Println("==============")
	foo(3, 0)
}
