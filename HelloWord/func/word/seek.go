package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("user.txt")
	//偏移量，相对位置
	//文件开始0 os.SEEK_SET(SEEK_SET=开始位)
	//当前位置1 os.SEEK_CUR
	//文件末尾2 os.SEEK_END
	bytes := make([]byte, 100)
	n, _ := file.Read(bytes)

	fmt.Println(n, bytes[:n])
	//也可以转成字符串形式
	fmt.Println(n, string(bytes[:n]))
	//偏移量，相对位置
	//文件的指针操作seek函数。
	file.Seek(0, 0)
	n, err := file.Read(bytes)
	fmt.Println(n, err, bytes[:n])

	file.Close()
}
