package main

import (
	"fmt"
	"net/http"
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

func main() {
	//启动服务还是要注册
	rpc.Register(&Calc{})
	rpc.HandleHTTP() //使用http请求
	////启动监听
	//listener,_ := net.Listen("tcp","8888")
	////循环接收
	//http.Serve(listener)
	http.ListenAndServe("0.0.0.0:8888", nil)

}
