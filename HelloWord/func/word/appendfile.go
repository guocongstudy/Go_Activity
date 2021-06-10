package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

//5.追加文件
func main() {
	path := "user.log"
	file, err := os.OpenFile(path, os.O_APPEND, os.ModePerm)

	if err == nil {
		file.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
		file.WriteString("\n")
		file.Close()
	}
	log.Print(time.Now().Unix())

}
