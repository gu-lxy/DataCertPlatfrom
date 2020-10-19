package main

import (
	"DataCertPlatfrom/blockchain"
	"DataCertPlatfrom/db_mysql"
	_ "DataCertPlatfrom/routers"
	"github.com/astaxie/beego"
)

func main() {
	block := blockchain.CreateGenesisBlock()
	blockchain.NewBlock(block0,Height+1, block0.Hash, []byte("a"))


	//连接数据库
	db_mysql.Connect()
	//静态资源文件映射设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()
}

