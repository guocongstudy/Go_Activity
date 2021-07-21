package main

import "fmt"

//匿名嵌入

type Sender2 interface {
	Send(msg string) error
}
type Reciver interface {
	Recive() (string, error)
}

type Client interface {
	//方法一：
	//Send(msg string)error
	//Reciver()(string,error)

	//方法二：
	//这就属于匿名嵌入
	Sender2
	Reciver
	Open() error
	Close() error
}
type MSNClient struct{}

func (c MSNClient) Send(msg string) error {
	fmt.Printf("发送信息：%d\n", msg)
	return nil
}
func (c MSNClient) Recive() (string, error) {
	fmt.Println("接收信息")
	return "", nil
}
func (c MSNClient) Open() error {
	return nil
}
func (c MSNClient) Close() error {
	return nil
}

func main() {

	msn := MSNClient{}
	var s Sender2 = msn
	var r Reciver = msn
	var c Client = msn

	s.Send("1")
	r.Recive()

	c.Open()
	defer c.Close()
	c.Send("2")
	c.Recive()

	var closer interface {
		Close() error
	}
	closer = msn
	closer.Close()
}
