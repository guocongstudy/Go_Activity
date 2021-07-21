package main

import (
	"encoding/json"
	"fmt"
)

/*
json.Marshal 序列化 内存->字符串/字节切片
json.Unmarshal 反序列化 字符串/字节切片 ->内存
*/
func main() {
	//写了一个切片
	names := []string{"郭聪", "李瑞", "小贝"}
	//定义一个映射
	users := []map[string]string{{"name": "郭聪", "address": "汉中", "number": "1231"}, {"name": "李瑞", "address": "汉中", "number": "13423"}}
	//格式化输出Marshal
	bytes, err := json.MarshalIndent(names, "", "\t")
	if err == nil {
		fmt.Println(bytes)
		//将字节转成字符串
		fmt.Println(string(bytes))
	}
	var names02 []string
	err = json.Unmarshal(bytes, &names02)
	fmt.Println(err)
	fmt.Println(names02)

	bytes, err = json.MarshalIndent(users, "", "\t")

	if err == nil {
		fmt.Println(string(bytes))
	}
	var user02 []map[string]string

	err = json.Unmarshal(bytes, &user02)
	fmt.Println(err)
	fmt.Println(user02)
	//用来判断字符串的格式是不是正确的
	fmt.Println(json.Valid([]byte("[]x")))
}
