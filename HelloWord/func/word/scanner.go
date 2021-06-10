package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("user.txt")
	if err == nil {
		defer file.Close()
		i := 0
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		for scanner.Scan() {
			//读取文件的一行
			fmt.Println(scanner.Text())
			i++
		}
	}
}
