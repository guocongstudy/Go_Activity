package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

type AddRequest2 struct {
	Left  int
	Right int
}

type AddResponse2 struct {
	Result int
}

func main() {
	client, _ := jsonrpc.Dial("tcp", "127.0.0.1:8888")
	req := AddRequest2{1, 100}
	resp := AddResponse2{}
	err := client.Call("calc.Add", req, &resp)
	fmt.Println(err, resp)
}
