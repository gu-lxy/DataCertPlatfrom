package models

import (
	"DataCertPlatform/db_mysql"
	"DataCertPlatform/utils"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
	Name     string `form:"name"` //名字
	Card     string `form:"card"` //身份证号
	Sex      string `form:"sex"`  //性别
}

/**
 * 该方法用于更新数据库中用户记录的实名认证信息
 */
func (u User) UpdateUser() (int64, error) {
	rs, err := db_mysql.Db.Exec("update user set name = ?, card = ?, sex = ? where phone = ?", u.Name, u.Card, u.Sex, u.Phone)
	if err != nil {
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/**
 * 将用户的信息保存到数据库中
 */
func (u User) AddUser() (int64, error) {
	//脱敏
	u.Password = utils.MD5HashString(u.Password) //把脱敏的密码的md5值重新赋值为密码进行存储

	rs, err := db_mysql.Db.Exec("insert into user(phone,password) values(?,?)",
		u.Phone, u.Password)
	//错误早发现早解决
	if err != nil { //保存数据遇到错误
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil { //保存数据遇到错误
		return -1, err
	}
	//id值代表的是此次数据操作影响的行数,id是一个整数int64类型
	return id, nil
}

/**
 * 查询用户信息
 */
func (u User) QueryUser() (*User, error) {

	//把脱敏的密码的md5值重新赋值为密码进行存储
	u.Password = utils.MD5HashString(u.Password)

	row := db_mysql.Db.QueryRow("select phone, name, card from user where phone = ? and password = ?",
		u.Phone, u.Password)

	err := row.Scan(&u.Phone, &u.Name, &u.Card)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (u User) QueryUserByPhone() (*User, error) {
	row := db_mysql.Db.QueryRow("select id, name, card, phone from user where phone = ?", u.Phone)
	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Card, &user.Phone)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
