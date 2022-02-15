package v1

import (
	"demo1/MyBlog/model"
	"demo1/MyBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(c *gin.Context) {
	profile, code := model.GetProfile(c)
	//fmt.Println("profile",profile)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    profile,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateProfile(c *gin.Context) {
	var req model.Profile
	_ = c.ShouldBindJSON(&req)

	profile, code := model.GetProfile(c)
	req.ID = profile.ID
	//fmt.Println("update profile",req)
	code = model.UpdateProfile(c, req.ID, &req)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
