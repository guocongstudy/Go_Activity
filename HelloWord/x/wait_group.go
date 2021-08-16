package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(n int) {
			defer wg.Done()
			// do something
			fmt.Printf("come to sub routine,arg =%d\n", n)

		}(i)
	}
	wg.Wait()
}
