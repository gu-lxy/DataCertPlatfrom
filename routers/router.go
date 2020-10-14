package routers

import (
	"DataCertPlatfrom/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router；路由
    beego.Router("/", &controllers.MainController{})
    //用户注册接口
    beego.Router("/register", &controllers.RegisterController{})
    //用户登录接口
    beego.Router("/login",&controllers.RegisterController{})
    //指用户上传的文件功能
    beego.Router("/upload",&controllers.UploadFilController{})
}
