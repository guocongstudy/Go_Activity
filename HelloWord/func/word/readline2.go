package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("user.txt")
	if err == nil {
		defer file.Close()
		reader := bufio.NewReader(file)

		for {
			//slice就是‘\n’的意思
			line, err := reader.ReadSlice('\n')
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Println(err, string(line))
			}

		}
	}
}
