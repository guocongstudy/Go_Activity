package main

import (
	"fmt"
	"testing"
)

func add5(x, y int) (sum int) {
	return x + y
}

func TestAdd(t *testing.T) { //不能少了这个* 报错找半天
	a := add5(5, 6)
	fmt.Println(a)
}
