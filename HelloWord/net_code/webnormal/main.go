package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*//提交数据常用规则
	//1.一个名字对应一个值
	//2.提交方式
	  GET
	  POST
	  编码方式：w-www-form-urlencoded
	          multipart/form-data
	第三方工具：curl,client
	  GET
	     参数常在URL
	  POST
	     body
	     w-www-form-urlencoded
	     multipart/form-data
	     application/json
	  DELETE
	     参数常在URL

	  PUT
	      body
	  HEAD
	      参数常在URL（取决于client工具/服务是否支持）
	*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.FormValue("x"))
		fmt.Println(r.PostFormValue("y"))
		r.FormFile("z")
	})
	http.ListenAndServe("0.0.0.0:8080", nil)

}
