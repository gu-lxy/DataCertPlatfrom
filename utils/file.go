package utils

import (
	"io"
	"mime/multipart"
	"os"
)

//保存一个文件
func SaveFile(fileName string, file multipart.File) (int64, error) {
	saveFile, err := os.OpenFile(fileName,os.O_CREATE|os.O_RDWR,777)
	if err != nil {
		return -1, err
	}

	length, err := io.Copy(saveFile, file)
	if err != nil {
		return -1, err
	}
	return length, nil
}
