package main

//定义一个请求
type Request struct {
	Left  int
	Right int
}

//定义一个响应
type Response struct {
	Result int
}

type Calc struct {
}

//定义一个响应方法
func (c *Calc) Sum(r Request, rp Response) error {
	rp.Result = r.Left + r.Right
	return nil
}
