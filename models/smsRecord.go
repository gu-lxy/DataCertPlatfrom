package models

import "DataCertPlatform/db_mysql"

type SmsRecord struct {
	BizId     string
	Phone     string
	Code      string
	Status    string
	Message   string
	TimeStamp int64
}

func (s SmsRecord) SaveSmsRecord(){
	db_mysql.Db.Exec("insert into sms_record(...")
}
