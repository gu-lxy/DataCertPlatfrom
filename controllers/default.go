package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

/**
 * 展示默认的首页：即用户注册页面
 */
func (c *MainController) Get() {
	//c.Ctx.WriteString("恭喜你，通过云服务器的公网ip访问到了我！")
	c.TplName = "register.html"
}
