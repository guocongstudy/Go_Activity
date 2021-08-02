package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	//解析模版 字符串
	text := "我的名字叫{{.}}"
	tpl, err := template.New("test").Parse(text)
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(os.Stdout, "KK")
	tpl.Execute(os.Stdout, 1)
	tpl.Execute(os.Stdout, true)
	tpl.Execute(os.Stdout, []string{"1", "2", "3"})

}

//4:17:09
