package utils

import (
	"io"
	"os"
)

//保存一个文件
func SaveFile(fileName string, reader io.Reader) (int64, error) {
	saveFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 777)
	if err != nil {
		return -1, err
	}

	length, err := io.Copy(saveFile, reader)
	if err != nil {
		return -1, err
	}
	return length, nil
}
