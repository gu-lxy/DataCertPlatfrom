package controllers

import (
	"DataCertPlatfrom/models"
	"DataCertPlatfrom/utils"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strings"
	"time"
)

type UploadFilController struct {
	beego.Controller
}

//该post方法用于处理用户在客户端提交的文件
func (u *UploadFilController) Post() {

	phone := u.Ctx.Request.PostFormValue("phone")
	title := u.Ctx.Request.PostFormValue("upload_title")

	fmt.Println("电子数据标签：", title)
	file, header, err := u.GetFile("gu")
	//defer File.Close()//空指针错误：invalid memory or nil pointer
	if err != nil {
		u.Ctx.WriteString("抱歉，文件解析失败，请重试")
		return
	}

	defer file.Close() //延迟执行

	//使用os包提供的方法保存文件
	//io.Copy(目标文件,数据源)
	saveFilePath := "static/upload" + header.Filename
	fmt.Println(saveFilePath)
	_, err = utils.SaveFile(saveFilePath, file)
	if err != nil {
		u.Ctx.WriteString("抱歉，文件数据认证失败")
		return
	}
	//2、计算文件的SHA256值
	fileHash, err := utils.SHA256HashReader(file)
	fmt.Println(fileHash)
	//查询用户id
	user1, err := models.User{Phone: phone}.QueryUserByPhone()
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败")
		return
	}

	//把上传的文件作为记录保存到数据库中
	//①  计算MD5值
	saveFile, err := os.Open(saveFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(saveFile)

	md5String, err := utils.MD5HashReader(file)
	if err != nil {
		fmt.Println("查询用户：", err.Error())
		u.Ctx.WriteString("抱歉，电子数据认证失败")
		return
	}

	record := models.UploadRecord{
		UserId:    user1.Id,
		FileName:  header.Filename,
		FileSize:  header.Size,
		FileCert:  md5String,
		FileTitle: title,
		CertTime:  time.Now().Unix(),
	}

	_, err = record.SaveRecord()
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证保存失败，请稍后再试!")
		return
	}
	//上传文件保存到数据库中成功
	records, err := models.QueryRecordsByUserId(user1.Id)
	if err != nil {
		u.Ctx.WriteString("抱歉, 获取电子数据列表失败, 请重新尝试!")
		return
	}
	u.Data["Records"] = records
	u.TplName = "list_record.html"

	u.Ctx.WriteString("恭喜，文件上传成功")
}

//该post方法用于处理用户在客户端提交的认证文件
func (u *UploadFilController) Post1() {
	title := u.Ctx.Request.PostFormValue("upload_title")

	//用户上传文件
	file, header, err := u.GetFile("gu")

	if err != nil { //解析客户端提交的数据
		u.Ctx.WriteString("抱歉，文件解析失败，请重试")
		return
	}
	defer file.Close() //延迟执行
	//获得上传文件
	fmt.Println("自定义的标题：", title)
	fmt.Println("上传的文件名称：", header.Filename)
	//
	fileNameSlice := strings.Split(header.Filename, ".")
	FileType := fileNameSlice[1]
	if FileType != "jpg" || FileType != "png" {
		u.Ctx.WriteString("抱歉，文件类型不符合，请上传正确的文件类型")
		return
	}

	isJpg := strings.HasSuffix(header.Filename, ".jpg")
	isPng := strings.HasSuffix(header.Filename, ".png")
	if !isJpg && !isPng {
		//文件类型不支持
		u.Ctx.WriteString("抱歉，文件类型不符合, 请上传符合格式的文件")
		return
	}
	fmt.Println("上传的文件大小：", header.Size)
	config := beego.AppConfig
	fileSize, err := config.Int64("file_size")
	if header.Size/1024 > fileSize {
		u.Ctx.WriteString("抱歉，文件大小超出范围，请上传大小合适的文件")
		return
	}
	fmt.Println("上传的文件大小:", file)

	//perm:permission 权限
	//权限的组成：a+b+c
	//a:文件所有者对文件的操作权限：读4、写2、执行1
	//b:文件所有者所在组的用户的操作权限，读4、写2、执行1
	//c:其他用户的操作权限，读4、写2、执行1
	saveDir := "static/upload"
	////打开文件
	_, err = os.Open(saveDir)
	if err != nil {
		//创建文件夹
		err = os.Mkdir(saveDir, 777)
		if err != nil {
			fmt.Println(err.Error())
			u.Ctx.WriteString("抱歉，文件认证遇到错误，请重试")
			return
		}
	}

	//文件名：文件路径+文件名+"."+文件扩展名
	saveName := "static/upload" + header.Filename
	fmt.Println("要保存的文件名：", saveName)
	//fromfile：文件
	//toFile：要保存的文件路径
	err = u.SaveToFile("file", saveName)
	if err != nil {
		u.Ctx.WriteString("抱歉，文件认证失败，请重试")
		return
	}
	fmt.Println("上传的文件：", file)
	u.Ctx.WriteString("已获取文件")
}
