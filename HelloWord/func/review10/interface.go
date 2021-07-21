package main

import "fmt"

type Sender interface {
	//定义了一个接口，type用于自定义类型，结构体中使用 接口里面有两个方法
	Send2(to string, msg string) error
	SendAll(tos []string, msg string) error
}

type EmailSender struct {
}

//定义send方法
func (s EmailSender) Send2(to, msg string) error {
	fmt.Println("发送邮件给:", to, "邮件内容是:", msg)
	return nil
}

func (s EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send2(to, msg)
	}
	return nil
}

type SmsSender struct {
}

func (s SmsSender) Send2(to, msg string) error {
	fmt.Println("发送信息给:", to, "消息内容是:", msg)
	return nil
}

func (s SmsSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		s.Send2(to, msg)
	}
	return nil
}

func do(sender SmsSender) {
	sender.Send2("领导", "工作日志")
}

func main() {
	//下面这一句还不清楚为什么这么操作
	//var sender Sender = EmailSender{}
	var sender Sender = SmsSender{}

	fmt.Printf("%T %v\n", sender, sender)

	sender.Send2("kk", "早上好！")
	sender.SendAll([]string{"祥哥", "烟灰"}, "中午好！")

}
