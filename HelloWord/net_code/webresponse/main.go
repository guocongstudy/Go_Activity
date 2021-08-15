package main

import (
	"fmt"
	"net/http"
)

var html = ` <!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8"/> 
        <title>我的第一个页面</title>
    </head>
    <body>
         <!-- 这里是注释，不在页面显示 -->
         我叫guocong
         <form action="" method="GET">
               <label>姓名</label><input  name="username" type="text"/>
               <label>密码</label><input  name="password" type="password" />
               <label>性别</label><input type="radio" name="sex" />
                        <input type="radio" name="sex" value="1"/> checked="checked"<lable>男<lable>
                        <input type="radio" name="sex" value="0"/><lable>女</lable>
    
               <input type="submit" value="登录" />
               <input type="button" value="button"/>
               <button>button按钮</button>
               <button type="submit">submit submit</button>
         </form>
    <label>爱好</label>
              <input type="checkbox" name="hobby" value="1"/><label>足球</label>
              <input type="checkbox" name="hobby" value="2"/><label>篮球</label>
              <input type="checkbox" name="hobby" value="3"/><label>乒乓球</label>
    <label></label>
    <label>备注</label>
        <textarea name ="remark">我叫...</textarea>
        
      <select>
             <option>开发</option>
             <option>测试</option>
      </select>
         <div>块1</div>
         <div>块2</div>
         <span>SPAN 1</span>
         <span>SPAN 2</span>
              <thead>
                  <tr>
                         <th>ID</th>
                         <th>姓名</th>
                         <th>位置</th>
                  </tr>
              </thead>
         <h1>一级标题</h1>
         <h2>一级标题</h2>
         <h3>一级标题</h3>
         <h4>一级标题</h4>
         <h5>一级标题</h5>
         <h6>一级标题</h6>
      <p>第一段。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。。</p>
      <p>第二段、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、</p>
      <a href="http://www.baidu.com" target="_black">百度</a>
    </body>
</html>
`

func main() {
	addr := ":8080"
	//动态的生成响应结果
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, html)
	})
	http.ListenAndServe(addr, nil)
}
