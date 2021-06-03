package main

//import (
//	"fmt"
//	"strconv"
//	"strings"
//)
//
////从命令行输入密码。并进行验证
////通过返回值告知验证成功还是失败
//
//const (
//	maxAuth  = 3
//	password = "123"
//)
//
//func auth() bool {
//	var input string
//	for i := 0; i <= maxAuth; i++ {
//		fmt.Println("请输入密码：")
//		fmt.Scan(&input)
//		if password == input {
//			return true
//		} else {
//			fmt.Print("密码错误\n")
//		}
//	}
//	return false
//}
//func printUser(pk int,users map[string]string){
//	fmt.Println("ID:",pk)
//	fmt.Println("名字:",users["name"])
//	fmt.Println("出生日期:",users["birthday"])
//	fmt.Println("联系方式:",users["tel"])
//	fmt.Println("联系地址:",users["addr"])
//	fmt.Println("备注:",users["desc"])
//}
//
//func query(users map[int]map[string]string) {
//
//	q := inputString("请输入查询内容：")
//	fmt.Println("==================================")
//	for k, v := range users {
//		if strings.Contains(v["name"], q) || strings.Contains(v["tel"], q) || strings.Contains(v["addr"], q) || strings.Contains(v["desc"], q) {
//			printUser(k,v)
//			fmt.Println("----------------------------------")
//		}
//	}
//}
//
//func getId(users map[int]map[string]string) {
//	var id int
//
//	for k, _ := range users {
//		if id < k {
//			id = k
//		}
//	}
//	return id + 1
//}
//
//func add(users map[int]map[string]string) {
//	id := getId(users)
//	users := inputUser()
//	users[id]=users
//	fmt.Println("[+]添加成功")
//}
//func modify(users map[int]map[string]string){
//	idString :=inputString("请输入修改用户ID：")
//	if id,err :=strconv.Atoi(idString);err ==nil{
//		if users,ok :=users[id];ok{
//			fmt.Println("将修改的用户信息：")
//			fmt.Println(users)
//			input:=inputString("确定修改(Y/N)?")
//			if input =="y"||input =="Y"{
//				users["name"] = inputString("请输入名字：")
//				users["birthday"] = inputString("请输入出生日期（2001-01-01）：")
//				users["tel"] = inputString("请输入联系方式：")
//				users["addr"] = inputString("请输入联系地址：")
//				users["desc"] = inputString("请输入备注：")
//				users[id] =users
//			}
//		}else{
//			fmt.Println("用户ID不存在")
//		}
//	}else {
//		fmt.Println("输入ID不正确")
//	}
//
//
//}
//func inputString(prompt string) string {
//	var input string
//	fmt.Print(prompt)
//	fmt.Scan(&input)
//	return strings.TrimSpace(input)
//}
//func inputUser(users map[string]string){
//	users :=map[string]string{}
//	users["name"] = inputString("请输入名字：")
//	users["birthday"] = inputString("请输入出生日期（2001-01-01）：")
//	users["tel"] = inputString("请输入联系方式：")
//	users["addr"] = inputString("请输入联系地址：")
//	users["desc"] = inputString("请输入备注：")
//	return users
//}
//
//func main() {
//	if !auth() {
//		//fmt.Println("密码错误，程序退出")
//		fmt.Printf("密码%d次错误，程序退出\n", maxAuth)
//		return
//	}
//	menu := `***********************************
//     1.查询
//     2.添加
//     3.修改
//     4.删除
//     5.退出
//******************************
//`
//	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "ID", "Name", "Age", "Tel", "Addr")
//	users := map[int]map[string]string{}
//	var id int
//	fmt.Println("欢迎进入gc的用户管理系统")
//END:
//	for {
//		fmt.Println(menu)
//
//		switch inputString("请输入指令：") {
//		case "1":
//			query(users)
//		case "2":
//			id++
//			add(users)
//		case "3":
//			modify(users)
//		case "4":
//
//		case "5":
//			break END
//
//		}
//	}
//}
