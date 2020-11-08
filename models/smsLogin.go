package models

type SmsLogin struct {
	Phone string `form:"phone"`
	Code  string `form:"code"`
}
