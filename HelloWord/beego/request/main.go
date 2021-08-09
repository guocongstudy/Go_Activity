package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginForm struct {
	UserName string `form:"username"` //这个标签就是指的是传入参数时用的形式
	Password string `form:"password"`
}
type RequestController struct {
	beego.Controller
}

//用户提交的数据
/*
Context =>c.Ctx
请求数据 c.Ctx
       c.Ctx.Request =>http.Request
             URL
                 ParseForm +Form
                 FormValue
             BODY
                 x-www-form-urlencoded
                     ParseForm + Form
                     ParseForm + Post Form
                     FormValue
                     PostFormValue
                 其他
                     Body
             c.Ctx.Input
Controller =>
   请求方法
      Get*
*/
func (c *RequestController) Header() {
	//请求控制器和动作
	fmt.Println(c.GetControllerAndAction())

	//请求头信息
	//请求行
	ctx := c.Ctx
	input := ctx.Input
	fmt.Println(input.Method())
	c.Ctx.Output.Body([]byte("header"))
}
func main() {
	beego.AutoRouter(&RequestController{})
	beego.Run()
}
