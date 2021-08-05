package main

import (
	"fmt"
	"net/http"
)

func main() {
	//現在还没有cookie信息
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("Cookie"))
		//Response Set-Cookie
		//设置cookie信息
		w.Header().Add("set-cookie", "xxx=xxxxx")
	})
	http.ListenAndServe(":9999", nil)
}
