package controllers

import (
	"DataCertPlatfrom/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}


//post方法处理用户登录请求
func (l *LoginController) Post() {
	//1、客户端提交的登录数据
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录信息解析失败")
		return
	}
	//2、根据解析到的数据，执行数据库查询解析操作
	u, err := user.QueryUser()
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录失败")
		return
	}


	//3、判断数据库查询结果

	//4、根据查询结果返回客户端相应的信息或者页面跳转
	l.Data["Phone"] = u.Phone//动态数据设置
	l.TplName = "home.html"


}
