package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	group := &sync.WaitGroup{}
	//信号量=2
	n := 2
	group.Add(n)

	for i := 1; i <= 2; i++ {
		go func() {
			for ch := 'A'; ch <= 'Z'; ch++ {
				fmt.Printf("%s:%c\n", "1", ch)
				runtime.Gosched()
			}
			group.Done()
		}()
	}
	group.Wait()

}
