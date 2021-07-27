package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//創建TCP鏈接
	addr := "127.0.0.1:999"
	url := "/net.go"
	client, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Fprintf(client, "GET %s HTTP/1.0\n", url)
	reader := bufio.NewReader(client)
	//响应行
	line, err := reader.ReadString('\n')
	fmt.Println("响应行：")
	fmt.Println("\t", line)
	fmt.Printf("响应头：")
	//响应头
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("错误信息:%d", err)
		}
		fmt.Printf(line)
	}

	client.Close()
}
