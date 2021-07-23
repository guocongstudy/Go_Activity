package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

/* Tcp/ip协议族
等待客户端链接，当客户端链接完成后，
给客户端发送一个当前时间，关闭客户端
*/

func main() {
	addr := "0.0.0.0:9999"
	//监听就用的listen
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer listener.Close()
	fmt.Println("Listen:", addr)

	input := bufio.NewScanner(os.Stdin)
	for {

		client, err := listener.Accept()
		if err == nil {
			reader := bufio.NewReader(client)
			writer := bufio.NewWriter(client)

			for {
				fmt.Print("请输入:")
				input.Scan()

				_, err := writer.WriteString(time.Now().Format("2021-03-06 15:05:06 "))
				writer.Flush()
				if err != nil {
					break
				}

				line, err := reader.ReadString('\n')
				if err != nil {
					break
				}
				fmt.Print("客户端：", strings.TrimSpace(line))
			}

			//客户端关闭
			client.Close()

		}

	}

}
