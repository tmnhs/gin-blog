package model

import (
	"context"
	"demo1/MyBlog/utils"
	"demo1/MyBlog/utils/errmsg"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
)
//使用七牛对象存储
var  AccessKey  = utils.AccessKey
var  SecretKey   =utils.SecretKey
var  Bucket      =utils.Bucket
var ImgUrl  = utils.QiniuServer

//todo 上传文件
func UploadFile(file multipart.File,fileSize int64)(string,int)  {
	putPolicy:= storage.PutPolicy{
		Scope: Bucket,
	}
	mac:=qbox.NewMac(AccessKey,SecretKey)
	upToken:=putPolicy.UploadToken(mac)

	cfg:=storage.Config{
		Zone: &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS: false,
	}

	putExtra:=storage.PutExtra{}

	formUploader:=storage.NewFormUploader(&cfg)
	ret:=storage.PutRet{}

	err:= formUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	if err != nil {
		return "",errmsg.ERROR
	}
	url:=ImgUrl+ret.Key
	return url,errmsg.SUCCESS
}
