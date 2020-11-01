package models

import (
	"bytes"
	"encoding/gob"
)

//该结构与体用于定义链上数据认证
type CertRecord struct {
	CerId []byte//认证id，本质是一个MD5值
	CertHash []byte//存证文件的hash值，本质是一个SAH256值
	CertName string//认证人姓名
	Phone string//联系方式
	CertCard string//身份证号
	FileName string//认证文件的名称
	FileSize int64//文件的大小
	CertTime int64//认证的时间
}

//序列化操作
func (c CertRecord) SeriaLiaze() ([]byte,error)  {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(c)
	return buff.Bytes(), err
}
//反序列化
func DeseriaLizeCertRecord(data []byte) (*CertRecord,error)  {
	var cerRecord *CertRecord
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&cerRecord)
	return cerRecord, err
}