package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	//对字符串操作的包 strings or bytes
	//内存字符串的操作
	//读，写
	//strings.Reader
	reader := strings.NewReader("123abcxyz")
	fmt.Println(reader.Len(), reader.Size())

	ctx := make([]byte, 5)
	n, err := reader.Read(ctx)
	fmt.Println(n, err, string(ctx[:n]))
	//strings.Builder
	reader.Seek(0, 0)
	fmt.Println(reader.Len(), reader.Size())
	//重置
	reader.Reset("xyz")
	reader.WriteTo(os.Stdout)

	//strings.Builder
	var builder strings.Builder
	builder.Write([]byte("ABC123"))
	builder.WriteString("xyz")
	fmt.Println(builder.String())

	//bytes
	//reader，buffer
	byteReader := bytes.NewReader([]byte("abc123abc"))
	byteReader.Read(ctx)
	fmt.Println(byteReader.Len(), byteReader.Size())
	fmt.Println(err, string(ctx[:n]))

	//Buffer 读/写
	buffer := bytes.NewBufferString("abc")
	buffer.WriteString("123")

	buffer.WriteString("xyz")
	//n,err =buffer.Read(ctx)
	//fmt.Println(err,string(ctx[:n]))

	fmt.Println(buffer.String())
}
