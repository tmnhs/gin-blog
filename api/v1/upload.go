package v1

import (
	"demo1/MyBlog/model"
	"demo1/MyBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

//todo 上传文件
func Upload(c *gin.Context)  {
	file ,fileHeadr,_:=c.Request.FormFile("file")

	fileSize:=fileHeadr.Size

	url,code:= model.UploadFile(file,fileSize)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
		"url":url,
	})
}