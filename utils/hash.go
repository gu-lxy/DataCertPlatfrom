package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

//对一个字符串MD5进行hash
func MD5HashString(data string) string{
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}

func MD5HashReader(reader io.Reader) (string, error)  {
	md5Hash := md5.New()
	readerBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	md5Hash.Write(readerBytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}


func SHA256HashReader(reader io.Reader) (string, error) {
	sha256 := sha256.New()
	readerBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	sha256.Write(readerBytes)
	hashBytes := sha256.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}