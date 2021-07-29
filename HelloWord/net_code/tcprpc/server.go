package main

import (
	"fmt"
	"net"
	"net/rpc"
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

//远程服务 Add(1,2)int
func main() {
	//req :=AddRequest{1,2}
	//resp :=AddResponse{}
	//calc :=Calc{}
	//calc.Add(req,&resp)
	//fmt.Println(resp.result)

	//rpc
	//注册
	rpc.Register(&Calc{})

	//启动服务
	listener, _ := net.Listen("tcp", "0.0.0.0:8888")
	rpc.Accept(listener)
	listener.Close()

	//2:43:32
}
