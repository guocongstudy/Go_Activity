package main

import (
	"fmt"
	"net/rpc"
)

//远程服务 Add(1,2) int
type AddRequest1 struct {
	Left  int
	Right int
}

//type AddResponse int
type AddResponse1 struct {
	Result int
}

func main() {
	//客户端是不用注册的，服务器端需要注册
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:8888")
	req := AddRequest1{3, 10}
	resp := AddResponse1{}
	err := client.Call("Calc.Add", req, &resp)
	fmt.Println(err, resp)
	client.Close()

}
