package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("user.txt")
	if err == nil {
		defer file.Close()
		writer := bufio.NewWriter(file)
		writer.WriteString("adafwfvvb")
		writer.Write([]byte("123423")) //主页byte是字节用（）不能用 {}
		writer.Flush()
	}
}
