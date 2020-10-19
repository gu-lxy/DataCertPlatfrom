package models

import (
	"DataCertPlatfrom/db_mysql"
	"DataCertPlatfrom/utils"
)

type User struct {
	Id       int `form:"id"`
	Phone    string `from:"phone"`
	Password string `from:"password"`
}
//将用户的信息保存到数据库中
func (u User) AddUser() (int64,error) {
	u.Password = utils.MD5HashString(u.Password)

	rs, err := db_mysql.Db.Exec("insert into user(phone,password) values(?.?)",
		u.Phone, u.Password)
	//错误早发现早决绝
	if err != nil {//保存数据遇到错误
		return -1,err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return -1,err
	}
	//id值代表的是此次数据操作影响的行数，id是一个整数int类型
	return id,nil
}

//查询用户信息
func (u User) QueryUser() (*User,error) {
	u.Password = utils.MD5HashString(u.Password)

	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ?",
	u.Phone, u.Password	)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,nil
}

func (u User) QueryUserByPhone() (*User, error) {
	row := db_mysql.Db.QueryRow("select id from user where phone = ?", u.Phone)
	var user User
	err := row.Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}