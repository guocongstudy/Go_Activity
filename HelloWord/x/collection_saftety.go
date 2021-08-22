package main

import (
	"fmt"
	"sync"
)

var arr = [10]int{}

type Student struct {
	Name string
	Age  int
}

var dict = make(map[int]string)

var student = Student{Name: "tom", Age: 18}

func main() {
	//两个协程同时修改并发的数组
	//偶数全为3奇数全为4
	//等待两个进程结束
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		//for i := 1; i < 10; i += 2 {
		//	arr[i] = 3
		//}
		//全局的修改结构体
		student.Name = "gucong"
	}()
	go func() {
		defer wg.Done()
		//for i := 0; i < 10; i += 2 {
		//	arr[i] = 4
		//}
		//全局的修改结构体
		student.Age = 28
	}()
	wg.Wait()
	//fmt.Println(arr)
	fmt.Printf("name=%s,age=%d\n", student.Name, student.Age)
}
