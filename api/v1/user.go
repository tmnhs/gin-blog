package v1

/**
用户处理模块的接口
*/

import (
	"demo1/MyBlog/model"
	"demo1/MyBlog/proto"
	"demo1/MyBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//todo 查询用户详细信息（包括文章）
func GetUserInfo(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetUserInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 添加用户
func AddUser(c *gin.Context) {
	var data proto.ReqAddUser
	var code int
	//绑定
	_ = c.ShouldBindJSON(&data)
	//msg,code= validator.Validate(&data)
	//if code!=errmsg.SUCCESS{
	//	c.JSON(http.StatusOK,gin.H{
	//		"status":code,
	//		"message":msg,
	//	})
	//	return
	//}
	//判断用户名是否存在
	code = model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		user := &model.User{
			Username: data.UserName,
			Password: data.Password,
			Email:    data.Email,
			Role:     data.Role,
			Status:   "Y",
		}
		result, _ := model.CreateUser(user)
		profile := &model.Profile{
			ID:    int(result.ID),
			Name:  result.Username,
			Email: result.Email,
		}
		//fmt.Println("create user:",result)
		code = model.CreateProfile(profile)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户

// todo 查询用户列表
func GetUsers(c *gin.Context) {
	//获取参数
	var req proto.ReqFindUser
	_ = c.ShouldBindJSON(&req)
	//fmt.Println(req)
	//如果pageSize等于0，则取消分页
	if req.PageSize == 0 {
		req.PageSize = -1
	}
	if req.PageNum == 0 {
		req.PageNum = -1
	}
	//调用函数查看所有用户
	data, total, err1 := model.GetUsers(req.IdOrName, req.PageSize, req.PageNum)
	code := errmsg.SUCCESS
	if err1 != nil {
		code = errmsg.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 编辑用户
func EditUser(c *gin.Context) {
	var req proto.ReqEditUser
	var code int
	//获取参数
	_ = c.ShouldBindJSON(&req)
	user, _ := model.GetUserInfo(req.Id)
	if user.Username != req.UserName {
		code = model.CheckUser(req.UserName)
		if code != errmsg.SUCCESS {
			c.JSON(code, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			})
			return
		}
	}
	code = model.EditUser(req.Id, &req)

	profileOld, _ := model.GetProfileById(req.Id)
	profile := &model.Profile{
		ID:     req.Id,
		Name:   req.UserName,
		Email:  req.Email,
		Desc:   profileOld.Desc,
		QqChat: profileOld.QqChat,
		WeChat: profileOld.WeChat,
		Weibo:  profileOld.Weibo,
		Img:    profileOld.Img,
		Avatar: profileOld.Avatar,
	}
	code = model.UpdateProfile(c, profile.ID, profile)
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
