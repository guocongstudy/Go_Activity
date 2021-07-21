package main

import (
	"fmt"
	"runtime"
	"sync"
)

func PrintfChras2(name string, group *sync.WaitGroup) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%s:%c\n", name, ch)
		//用完把CPU让出去，谁抢到算谁的
		runtime.Gosched()
	}
	group.Done()
}

//要改数据，一般都是指针接收的方法
func main() {
	group := &sync.WaitGroup{}
	group.Add(2)
	//启动一个例程，就是go + 一个函数
	go PrintfChras2("1", group)
	go PrintfChras2("2", group)

	PrintfChras2("3", group)
	//加上group.Wait()才让你运行过程能够看得见
	group.Wait()
}
