package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//常用的就是这几种
	//base64.RawStdEncoding
	//base64.URLEncoding
	//base64.RawURLEncoding
	//将字符串进行加密
	x := base64.StdEncoding.EncodeToString([]byte("guocong"))
	fmt.Println(x)
	//加密后进行解密
	bytes, err := base64.StdEncoding.DecodeString(x)
	fmt.Println(string(bytes), err)
}
