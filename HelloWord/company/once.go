package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	for i := 0; i < 10; i++ {
		//代码只想执行一次
		once.Do(func() {
			fmt.Println(i)
		})
	}
}
