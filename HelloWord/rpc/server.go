package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

func main() {
	//注册rpc服务
	rpc.Register(&Calc{})

	server, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer server.Close()
	for {
		client, err := server.Accept()
		if err == nil {
			jsonrpc.ServeConn(client)
		}
	}

}
