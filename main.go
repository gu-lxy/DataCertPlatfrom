package main

import (
	_ "DataCertPlatform/routers"
	"DataCertPlatform/blockchain"
	"github.com/astaxie/beego"
	"DataCertPlatform/db_mysql"
)

func main() {

	//先准备一条区块链
	blockchain.NewBlockChain()

	//连接数据库
	db_mysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")

	beego.Run() //阻塞
	//http.ListenAndServe(":8080")
}
