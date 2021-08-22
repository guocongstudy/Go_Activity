package main

import (
	"sync"
	"time"
)

//資源競爭
var n int

func inc() {
	n++
}

var lock sync.RWMutex

//func inc2(){
//	atomic.AddInt32(&n,1)
//}
func inc3() {
	//加锁和写锁之间可能有多行代码

	lock.Lock()
	n++
	a := n
	b := a + 1
	n = b
	lock.Unlock()
}

//读写锁的问题
func main() {
	go inc()
	go inc()
	go inc()
	time.Sleep(1 * time.Second)
}
