package models

type User struct {
	Id int `form:"id"`
	Phone string `from:"phone"`
	Password string `from:"password"`
}
