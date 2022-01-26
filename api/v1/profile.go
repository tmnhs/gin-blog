package v1

import (
	"demo1/MyBlog/model"
	"demo1/MyBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProfile(c * gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	profile,code:=model.GetProfile(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":profile,
		"message":errmsg.GetErrMsg(code),
	})
}


func UpdateProfile(c * gin.Context)  {
	var profile model.Profile
	id,_:=strconv.Atoi(c.Param("id"))
	_=c.ShouldBindJSON(&profile)

	code := model.UpdateProfile(id,&profile)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}