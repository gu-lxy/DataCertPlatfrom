package controllers

import (
	beego "beego-develop"
	"fmt"
	"strings"
)

type UploadFilController struct {
	beego.Controller
}
func (u *UploadFilController) Post(){
	title := u.Ctx.Request.PostFormValue("upload_title")

	//用户上传文件
	File, header, err := u.GetFile("gu")
	if err != nil {
		u.Ctx.WriteString("抱歉，文件解析失败，请重试")
	}
	//获得上传文件
	fmt.Println("自定义的标题：",title)
	fmt.Println("上传的文件名称：",header.Filename)
	//
	fileNameSlice := strings.Split(header.Filename,".")
	FileType := fileNameSlice[1]
	if FileType != "jpg" || FileType != "png" {
		u.Ctx.WriteString("抱歉，文件类型不符合，请上传正确的文件类型")
		return
	}

	isJpg := strings.HasSuffix(header.Filename,".jpg")
	isPng := strings.HasSuffix(header.Filename,".png")
	if !isJpg && !isPng{
		//文件类型不支持
		u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
		return
	}
	fmt.Println("上传的文件大小：",header.Size)
	config := beego.AppConfig
	fileSize, err := config.Int64("file_size")
	if header.Size / 1024 > fileSize {
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传大小合适的文件")
		return
	}
	fmt.Println(File)
	u.Ctx.WriteString("已获取文件")
}