package controllers

import "github.com/astaxie/beego"

type LoginSmsController struct {
	beego.Controller
}

/**
 * 浏览器发起的短信验证码的登录请求
 */
func (l *LoginSmsController) Get() {
	l.TplName = "login_sms.html"
}

/**
 * 短信验证码登录的功能
 */
func (l *LoginSmsController) Post() {

}
