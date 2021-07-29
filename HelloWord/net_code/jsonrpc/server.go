package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type AddRequest struct {
	left  int
	right int
}

//type AddResponse int
type AddResponse struct {
	result int
}
type Calc struct{}

//参数1：请求对象(可以是指针，也可以是值)
//参数2：相应对象(可以是指针)
func (c *Calc) Add(rq AddRequest, resp *AddResponse) error {
	fmt.Println("calc.add")
	resp.result = rq.left + rq.right
	return nil
}

type AAA struct {
}

func (a *AAA) Add(req AddRequest, resp *AddResponse) error {
	fmt.Println("aaa.Add")
	return nil
}

func main() {
	//注册
	rpc.Register(&Calc{})
	listener, _ := net.Listen("tcp", "0.0.0.0:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go jsonrpc.ServeConn(conn)
	}
	listener.Close()
}
