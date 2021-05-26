package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%T\n", now)
	fmt.Printf("%v\n", now)
	//2000/01/02 08:10:01
	//2006年
	//01月
	//02天
	//24进制小时 15
	//分钟 04
	//秒 05
	fmt.Println(now.Format("2006/01/02 12:06:03"))
	fmt.Println(now.Format("2006-01-02 15:05:07"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:05:06"))

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//将字符串格式转成时间格式
	t, err := time.Parse("2006-01-02 15:04:05", "2006-01-02 02:08:03")
	fmt.Println(t, err)

	//所有Unix开始的时间
	t = time.Unix(0, 0)
	fmt.Println(t)

	d := now.Sub(t)
	fmt.Printf("%T,%v", d, d)

	fmt.Println(time.Now())
	time.Sleep(time.Second * 5)
	fmt.Println(time.Now())

	u := now.Add(3 * time.Hour)
	fmt.Println(u)

	d, err = time.ParseDuration("3h2m4s")
	fmt.Println(err, d)
	//转换成多少小时
	fmt.Println(d.Hours())
	//转换成多少分钟
	fmt.Println(d.Minutes())
	//转化成多少分钟
	fmt.Println(d.Seconds())
}
