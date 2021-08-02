package main

import "fmt"

//结构体能不能用等于还是不等于

type IPAndStatusCode struct {
	ip   string
	code int
}

func main() {
	v1 := IPAndStatusCode{"1.1.1.1", 200}
	v2 := IPAndStatusCode{"1.1.1.1", 200}
	if v1 == v2 {
		fmt.Println("两个结构体相等")
	} else {
		fmt.Println("两个结构体不等")
	}
}
