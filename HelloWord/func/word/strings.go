package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("adawdwfcvf")
	//byte就是字节的意思，定义byte类型，[]byte就是字节数组。就是一个数组，里面的每一个元素都是字符，字符又
	//跟字节划等号，所以字符串和字节数组之间可以相互转化。
	/*
	   1.string 转为[]byte
	   var str string ="test"
	   var data []byte = []byte(str
	   )
	*/
	//bytes := make([]byte, 3)
	//for {
	//	n, err := reader.Read(bytes)
	//	if err != nil {
	//		if err != io.EOF {
	//			fmt.Println(err)
	//		}
	//		break
	//	} else {
	//		fmt.Println(n, bytes[:n])
	//	}
	//}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	var builder strings.Builder

	builder.Write([]byte("abc"))
	builder.WriteString("sadaegv!")

	fmt.Println(builder.String())

}
