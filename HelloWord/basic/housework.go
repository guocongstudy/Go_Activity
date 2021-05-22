package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*1.打印乘法口诀
for ,if ,fmt.Printf, fmt.Println
  2.猜数字游戏
for if continue/break ...
    a.生成一个[0,100)随机数
    b.让用户重复猜5次(从命令行输入)
       猜得太大 =>太大,你还有几次猜词机会
       猜的太小 =>太小,你还有几次猜词机会
       猜的相等 =>猜中

       5次都没猜中 = >退出，并提示太笨了
  挑战：
  当5次都没猜中提示：太笨了，重新开始猜数游戏（重新生成随机数，重新开始5次计数）
*/

//打印99乘法表
//func main() {
//var i int
//var j int
//	for i=1;i<=9;i++{
//		//for j=1;j<=9;j++{ //出错原因
//		for j =1;j<=i;j++{
//			fmt.Printf("%d*%d=%d",i,j,i*j)
//			fmt.Printf(" ")
//			//fmt.Printf("%d *%d = %d\t",i,j,i*j) //替换上面两句，还是不行
//		}
//		fmt.Println()
//	}
//}
//问题：少了一句j<=i
//自己写的问题就是不会写怎么将程序弄成阶梯形状，这样输出就只能连续的。

//正确写法
//func main(){
//	for line:=1;line<=9;line++{
//		for i:=1;i<=line;i++{
//			fmt.Print(i,"*",line,"=",i*line,"\t")
//		}
//		fmt.Println()//打印换行
//	}
//
//}

/*
复习
1.fmt.Printf 需要指定模板 %s=字符串占位，%d=int类型，%f=浮点类型，%t=?,%p=指针类型，%v=不知道用那种数据类型，就用%v
%T= 查看某一个数据类型
2.类型名：uint，rune，uintptr
3.特殊字符：\n,\t,\r 这些什么意思？
4.for range 通常遍历的是“字符串，数组，切片，映射，管道”
*/

//作业2

//func main(){
//	//设置随机数种子
//	rand.Seed(time.Now().Unix())
//
//	//生成随机数0-100
//	//guess:=rand.Int()%100
//	guess :=rand.Intn(100)
//	const maxGuessNum = 5
//
//	for i:=0;i<maxGuessNum;i++{
//		var input int
//		fmt.Print("请输入你猜的值：")
//		fmt.Scan(&input)
//		if guess ==input{
//			fmt.Printf("你猜对了，太聪明了")
//		}else {
//			if guess <=input{
//				fmt.Printf("你猜的数太小了")
//			}else {
//				fmt.Printf("你猜的数太大了")
//			}
//		}
//		fmt.Printf("请重新输入：")
//	}
//}

//正确方法
func main() {
	//设置随机数种子
	rand.Seed(time.Now().Unix())
	var input int
	//生成随机数0-100
	//guess:=rand.Int()%100
	guess := rand.Intn(100)
	const maxGuessNum = 5

	for i:=0;i<maxGuessNum;i++{
		fmt.Print("请输入你要猜的数：")
		fmt.Scan(&input)

		if guess ==input{
			fmt.Printf("太聪明了，你第%d次猜对",i+1)
			break
		}else if guess <=input{
			fmt.Printf("你猜的太大了，你还有%d次机会\n",maxGuessNum-i-1)
		}else if guess >=input{
			fmt.Printf("你猜的太小了，还有%d次机会\n",maxGuessNum-i-1)
		}
		//当最后一次机会用完还没有猜对
		if i==maxGuessNum-1{
			fmt.Printf("你太笨了，5次都没有猜对")
		}
		//挑战
		var text string
		fmt.Printf("重新开始游戏吗？(y/n)")
		fmt.Scan(&text)
		if text !="y"{
			break
		}
	}

}
