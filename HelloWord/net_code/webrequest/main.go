package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	addr := "0.0.0.0:8888"
	//处理器函数，又写了一个匿名函数

	//http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	//
	//})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(strings.Repeat("-", 30))
		//1.提交方式
		/*
			  在url中传递数据
			样式：url?argname1=argvalue1&argname2=argvalue2
		*/
		r.ParseForm()                //解析参数
		fmt.Println(r.Form)          //接收的参数类型都是string
		fmt.Println(r.Form.Get("x")) //只返回第一个参数
		fmt.Println(r.Form["x"])

		//2.通过body提交数据格式
		/*curl -d “xxxxxxx”
				  body中数据格式
				  3种
				  application/x-www-form-urlencoded a=b&c=d
				  application/json

				其他类型
				  multipart/form-data
		          application/xml
				  自己定义

		*/

		//提交数据的编码格式是什么？
		//Request Header:Content-type
		fmt.Println(r.Header)
		//Header获取Content-type
		//json =>jsonParser（如果是json格式就用jsonParser解析）
		//x-www-form-urlencoded:ParseForm
		//xml =>xmlParser （如果是其他格式，就用其他格式去解析就行。）

		/*
		   a. application/x-www-form-urlencoded a=b&c=d
		   Form 可以获取URL中的参数也可以获取body中参数
		   PostForm 只可以获取body中的参数

		*/
		//对于自定义类型需要获取body原始数据
		//使用特定解码器
		//io.Copy(os.Stdout,r.Body)
		ctx, _ := ioutil.ReadAll(r.Body)
		var j map[string]interface{}
		json.Unmarshal(ctx, &j)
		fmt.Printf("%#v\n", j)

		fmt.Fprintf(w, time.Now().Format("2021-7-29 15:04:05"))

		////请求行 method Url 协议
		//fmt.Println("method:", r.Method)
		//fmt.Println("url:", r.URL)
		//fmt.Println("protocol:", r.Proto)
		////请求头
		//header := r.Header
		//fmt.Println(header.Get("User-Agent"))
		//fmt.Println(r.Header)
		//fmt.Fprint(w, time.Now().Format("2006-01-02 15:07:3"))
		////请求体
		////fmt.Fprint(w, time.Now().Format("2021-07-29 10:54:59"))
		////用copy把他拷贝到标准输出上
		//fmt.Println("body:")
		//io.Copy(os.Stdout,r.Body)

	})
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		//对于自定义类型需要获取body原始数据
		//使用特定解码器
		//io.Copy(os.Stdout,r.Body)
		//将一个r.body读取到字节切片中

	})
	http.HandleFunc("/file/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("------------------------")
		r.ParseMultipartForm(1024 * 1024) //接收文件过程中最大使用的内存
		//url ?
		//body k=v
		//body 文件内容
		fmt.Println(r.MultipartForm.File)
		fmt.Println(r.MultipartForm.Value)
		fmt.Println(r.Form)
		fmt.Println(r.PostForm)
		if fileHeaders, ok := r.MultipartForm.File["x"]; ok {
			for _, fileHeader := range fileHeaders {
				fmt.Println(fileHeader.Filename, fileHeader.Size)
				nfile, _ := os.Create("./file" + fileHeader.Filename)
				file, _ := fileHeader.Open()
				io.Copy(os.Stdout, file)
				file.Close()
				nfile.Close()
			}
		}

	})
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err)
}
