package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	//获取连接，定义请求对象
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:9999")
	if err == nil {
		request := Request{5, 10}
		var response Response

		err := client.Call("Calc.Sum", request, response)
		if err == nil {
			fmt.Println(response.Result)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}
