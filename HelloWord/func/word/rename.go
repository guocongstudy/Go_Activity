package main

import (
	"fmt"
	"os"
)

func main() {
	//给文件更换名字，多运行几次，第一次看可能没换好。
	os.Rename("user.log", "user2.log")
	//os.Remove("user2.log")

	file, err := os.Open("xxx")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		}
	} else {
		file.Close()
	}

	for _, path := range []string{"xxx", "reader.go", "user.txt"} {
		fileInfo, err := os.Stat(path)

		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("文件不存在")
			}
		} else {
			fmt.Println(fileInfo.Name())
			fmt.Println(fileInfo.IsDir())
			fmt.Println(fileInfo.Size())
			fmt.Println(fileInfo.ModTime())

			if fileInfo.IsDir() {
				dirfile, err := os.Open(path)
				if err == nil {
					defer dirfile.Close()

					names, _ := dirfile.Readdirnames(-1)
					for _, name := range names {
						fmt.Println(name)
					}
				}
			}
		}
	}
}
