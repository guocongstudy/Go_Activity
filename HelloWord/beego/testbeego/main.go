package main

//
//import (
//	"fmt"
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
//)
//
////路由控制器
//type HomeController struct {
//	beego.Controller
//}
//
//func (c *HomeController) Get() {
//	fmt.Println("HomeController.Get")
//	//没有指定响应:views/homecontroller/get.tpl
//	c.Ctx.Output.Body([]byte("HomeController.Get"))
//
//}
//
//func (c *HomeController) Post() {
//	fmt.Println("HomeController.Get")
//	//没有指定响应:views/homecontroller/get.tpl
//	c.Ctx.Output.Body([]byte("HomeController.Post"))
//
//}
//type AuthController struct {
//	beego.Controller
//}
//
//func (c *AuthController) Login(){
//	c.ctx.Output.Body([]byte("login"))
//}
//
//func (c *AuthController) Logout(){
//	c.ctx.Output.Bdoy([]byte("logout"))
//}
//
//
//func main() {
//	//2.0.0 beta
//	//1.12
//	//定义处理器
//	//绑定URL和处理器的功能 路由
//	//基本路由
//	//为URL 绑定某个请求方法的 路由函数
//	//为根目录的GET办法绑定函数
//	//POST,DELETE,HEAD,PUT,OPTIONS,PATCH
//	//后面这个func函数就是路由函数
//	beego.Get("/", func(ctx *context.Context) {
//		//和response 函数类似
//		ctx.Output.Body([]byte("hi,beego"))
//	})
//
//	beego.Post("/", func(ctx *context.Context) {
//		ctx.Output.Body([]byte("post"))
//	})
//
//	beego.Delete("/", func(ctx *context.Context) {
//		ctx.Output.Body([]byte("delete"))
//	})
//	//绑定所有的请求方法叫Any
//	beego.Any("/", func(ctx *context.Context) {
//		ctx.Output.Body([]byte("any"))
//	})
//
//	beego.Any(`/file/*.*`, func(ctx *context.Context) {
//		ctx.Output.Body([]byte("file:" + ctx.Input.Param("file")))
//	})
//
//	beego.Router("/home/", &HomeController{})
//	//提交参数
//	//？a=b
//	//body
//	//delete/1/
//	//正则路由 /delete/数值/ =>数值解析到参数中
//	//(\d{1,8})这个是限制了这个id的格式只能是1位到9位之间
//	beego.Any("/delete/:id(\\d{1,8})/", func(ctx *context.Context) {
//		ctx.Output.Body([]byte("delete" + ctx.Input.Param(":id")))
//	})
//	beego.Any(`get/:id:int`, func(ctx *context.Context) {
//		ctx.Output.Body([]byte("get:" + ctx.Input.Param(":id")))
//	})
//
//	//匹配uer开头的
//	beego.Any(`/user/*`, func(ctx *context.Context) {
//		ctx.Output.Body([]byte("user" + ctx.Input.Param(":splat")))
//	})
//	//注解路由
//	//启动服务
//	beego.Run()
//}
