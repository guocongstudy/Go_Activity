package main

import (
	"fmt"
	"net/http"
	"time"
)

//gow 安装上以后可以在win上使用一些linux的指令
func test02(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time:%d", time.Now().Unix())
}

type Test03 struct{}

func (t Test03) ServeHttP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time:%s", time.Now().Format("2021-01-02 15:04:05"))
}

func main() {
	//定义处理器函数
	http.HandleFunc("/test01/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi kk"))
	})

	http.HandleFunc("/test02", test02)

	err := http.ListenAndServe("0.0.0.0:9999", nil)
	if err != nil {
		fmt.Println(err)
	}
}
