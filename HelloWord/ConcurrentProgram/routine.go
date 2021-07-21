package main

import (
	"fmt"
	"runtime"
	"time"
)

//day11
//2. 例程

func PrintfChars() {
	//打印从大写字母A到大写字母Z

	var c = []string{"A", "B", "C"}
	for _, b := range c {
		fmt.Println(b)
	}
}

func PrintfChras1(name string) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", name, ch)
		//用完把CPU让出去，谁抢到算谁的
		runtime.Gosched()
	}
}

func main() {
	go PrintfChras1("1")
	go PrintfChras1("2")
	//1.在未sleep之前，看着只是执行了3,1和2未执行
	time.Sleep(3 * time.Second)
	//没有加go的是主例程，主例程执行完，其他没有完也就结束了
	PrintfChras1("3")
}
