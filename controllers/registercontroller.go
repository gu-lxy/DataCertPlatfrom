package controllers

import (
	"DataCertPlatfrom/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

//该方法用于处理用户注册的逻辑
func (r *RegisterController) Post(){
	r.TplName = "register.html"

	fmt.Println("hello register")
	//1、解析用户提交的请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据解析失败，请重试")
		return
	}
	//2、将解析到底数据保存到数据库中
	//3、将处理的结果返回给客户浏览器
	  //3.1 如果成功，跳转登录页面
	  r.TplName = "login.html"
	  //3.1 如果失败，提示错误信息
}
