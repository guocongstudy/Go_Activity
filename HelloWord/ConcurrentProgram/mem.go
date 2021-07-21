package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int
	group := &sync.WaitGroup{}
	//定义一个锁
	lock := &sync.Mutex{}
	incr := func() {

		defer group.Done()
		for i := 0; i <= 100; i++ {
			//加锁的用处在于，将这个执行完在执行下面的，不能一会做这个，一会做那个
			lock.Lock() //加锁（互斥锁）
			counter++
			runtime.Gosched()
			lock.Unlock() //释放锁
		}
	}
	decr := func() {
		defer group.Done()
		for i := 0; i <= 100; i++ {
			lock.Lock()
			counter--
			lock.Unlock()
			runtime.Gosched()
		}
	}
	for i := 0; i < 10; i++ {
		group.Add(2)
		go incr()
		go decr()
	}
	group.Wait()
	fmt.Println(counter)
}
