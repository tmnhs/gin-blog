package v1

/**
用户处理模块的接口
*/

import (
	"demo1/MyBlog/model"
	"demo1/MyBlog/utils/errmsg"
	"demo1/MyBlog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//todo 查询用户详细信息（包括文章）
func GetUserInfo(c *gin.Context)  {

	id, _ := strconv.Atoi(c.Param("id"))

	 data,code:=model.GetUserInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
//todo 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	var code int
	//绑定
	_ = c.ShouldBind(&data)
	msg,code= validator.Validate(&data)
	if code!=errmsg.SUCCESS{
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":msg,
		})
		return
	}
	//判断用户名是否存在
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		code = model.CreateUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户

// todo 查询用户列表
func GetUsers(c *gin.Context) {
	//获取参数
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	//如果pageSize等于0，则取消分页
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	//调用函数查看所有用户
	data ,total:= model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":total,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	//获取参数
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		code = model.EditUser(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 删除用户
func DeleteUser(c *gin.Context) {
	//获取参数
	id, _ := strconv.Atoi(c.Param("id"))
	//调用删除的函数
	code := model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}


