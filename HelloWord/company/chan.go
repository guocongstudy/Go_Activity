package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	//延迟多少秒后执行，就可以用after
	channel := time.After(3 * time.Second)
	fmt.Println(<-channel)
	//after一般只读一次，第二次读的时候就死锁了
	//fmt.Println(<-channel)
	//after只打印一次，ticker可以打印多次

	//每隔多少秒获取一次，产生一个管道消息
	ticker := time.Tick(3 * time.Second)
	fmt.Println(<-ticker)
	fmt.Println(<-ticker)
	fmt.Println(<-ticker)

	for now := range ticker {
		fmt.Println(now)
	}
}
