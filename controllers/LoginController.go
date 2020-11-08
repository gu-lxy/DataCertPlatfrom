package controllers

import (
	"github.com/astaxie/beego"
	"DataCertPlatform/models"
	"fmt"
	"strings"
)

type LoginController struct {
	beego.Controller
}

/**
 * 直接跳转展示用户登录页面
 */
func (l *LoginController) Get() {
	l.TplName = "login.html"
}

/**
 * post方法处理用户的登录请求
 */
func (l *LoginController) Post() {
	//1、解析客户端用户提交的登录数据
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录信息解析失败，请重试")
		return
	}

	//2、根据解析到的数据,执行数据库查询操作
	u, err := user.QueryUser()

	//3、判断数据库查询结果
	if err != nil {
		// sql: no rows in result set(集合），结果集中没有数据
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录失败，请重试")
		return
	}

	//3.1 增加逻辑: 判断用户是否已实名认证，如果没有实名认证，则跳转到认证页面，执行认证业务
	if strings.TrimSpace(u.Name) == "" || strings.TrimSpace(u.Card) == "" { //两者有其一，即为没有进行实名认证
		l.Data["Phone"] = user.Phone
		l.TplName = "user_kyc.html"
		return
	}

	//4、根据查询结果返回客户端相应的信息或者页面跳转
	l.Data["Phone"] = u.Phone //动态数据设置
	l.TplName = "home.html"   // 文件上传界面{{.Phone}
}
