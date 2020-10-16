package db_mysql

import (
	beego "beego-develop"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//Db数据库连接池
var Db *sql.DB

//数据库的连接
func Connect() {
	//项目配置
	config := beego.AppConfig
	driver := config.String("db_driver")//数据库驱动
	dbUser := config.String("db_user")//数据库用户名
	dbPasssword := config.String("db_password")//密码
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(driver,dbUser,dbPasssword,dbIp,dbName)
	//连接数据库
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	fmt.Println(connUrl)
	db,err:=sql.Open(driver,connUrl)
	if err != nil {
		panic("数据库连接失败。请重新尝试")
	}
	Db = db
}

