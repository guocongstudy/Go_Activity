package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//验证能否转成json的函数
	fmt.Println(json.Valid([]byte("[]")))

}
