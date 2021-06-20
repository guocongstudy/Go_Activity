package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//os.Stdout,标准输入输出
	fmt.Println("xxx")
	os.Stdout.Write([]byte("xxx"))

	bytes := make([]byte, 5)
	n, err := os.Stdin.Read(bytes)
	fmt.Println(n, err, bytes)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	fmt.Println(scanner.Text())

	fmt.Fprintf(os.Stdout, "%T", 1)
}
