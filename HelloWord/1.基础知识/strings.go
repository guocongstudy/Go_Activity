package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(strings.Compare("abc","bac"))
	fmt.Println(strings.Contains("abc","bc"))
	fmt.Println(strings.ContainsAny("abc","a"))
	fmt.Println(strings.Count("adadwdawd","ad"))
	fmt.Printf("%q\n",strings.Fields("abc def\neede\rdede\fde\vded"))//按空白字符分隔
    fmt.Println(strings.HasPrefix("abcd","ab"))
	fmt.Println(strings.HasSuffix("abcdef","defd"))
	fmt.Println(strings.Index("adadwd","def"))
	fmt.Println(strings.LastIndex("dwdwdaf","dw"))

	fmt.Println(strings.Split("abcdes;adc;ccd",";")) //按分号去分隔
	fmt.Println(strings.Join([]string{"qsQS","SQS","sqs"},":"))

	fmt.Println(strings.Replace("adcadcadaca","ad","xxx",1))//1指的是切换一次
	fmt.Println(strings.Replace("adadaddadad","ad","xxx",-1))//-1就是切换全部
	fmt.Println(strings.Title("hi,kk"))

	fmt.Println(strings.Trim("sdadadwa","sd"))
	fmt.Println(strings.TrimSpace("adcsd xxx \n \r"))


}
