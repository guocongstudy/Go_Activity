package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintChars(name int, channel chan int) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		fmt.Printf("%d:%c\n", name, ch)
		runtime.Gosched()
	}
	channel <- name
	fmt.Println("写入:", name)
}
func main() {
	var channel chan int = make(chan int)

	for i := 0; i < 10; i++ {
		go PrintChars(i, channel)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("读取：", i)
	}
	fmt.Println("over")
	time.Sleep(1 * time.Second)
}
