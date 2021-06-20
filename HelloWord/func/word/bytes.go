package main

import (
	"bytes"
)

func main() {
	buffer := bytes.NewBufferString("guocong")
	//在buffer中写入一个字节数组。
	buffer.Write([]byte("123"))
	buffer.WriteString("!#4")

	bytes := make([]byte, 2)
	buffer.Read(bytes)
}
