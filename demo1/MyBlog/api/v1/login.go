package v1

import (
	"demo1/MyBlog/middleware"
	"demo1/MyBlog/model"
	"demo1/MyBlog/utils"
	"demo1/MyBlog/utils/errmsg"
	"demo1/MyBlog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

//todo 注册
func Register(c *gin.Context) {
	var user model.User
	c.ShouldBindJSON(&user)
	//对传过来的数据进行校验
	msg,code:= validator.Validate(&user)
	//校验失败，返回错误信息（注册失败）
	if code!=errmsg.SUCCESS{
		c.JSON(http.StatusOK,gin.H{
			"code":code,
			"message":msg,
		})
		return
	}
	//判断用户名是否存在
	code = model.CheckUser(user.Username)
	if code==errmsg.SUCCESS{
		//用户名不存在，可以注册
		//设置激活码，唯一字符串
		user.Code=utils.CreateUUID()
		//设置激活状态
		user.Status="N"
		//激活邮件发送,邮件正文

		content:="<a href = 'http://localhost/api/v1/active?code="+user.Code+"'>点我激活</a>"

		model.SendEmail(user.Email,"激活邮件",content)

		//将注册的用户保存到数据库
		code= model.CreateUser(&user)
	}
	c.JSON(http.StatusOK,gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})

}

//todo 邮件激活
func ActiveEmail(c *gin.Context) {
	var code int
	//获取参数（激活码）
	vCode:=c.Query("code")
	if len(vCode)==0 {
		code=errmsg.ERROR_EMAIL_CODE_NOT_EXIST
		c.JSON(http.StatusOK,gin.H{
			"code":code,
			"message":errmsg.GetErrMsg(code),
		})
		return
	}
	code=model.UpdateUserStatus(vCode)

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"message":errmsg.GetErrMsg(code),
	})
}

//todo 登录
func Login(c *gin.Context) {
	var data model.User
	c.ShouldBindJSON(&data)
	//fmt.Println(data)
	var token string
	var code int
	code = model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(data.Username)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

//todo 发送邮件
func SendEmailForCode(c *gin.Context)  {
	to:=c.Query("email")
	code:=utils.CreateVcode()
	content:="<h2>欢迎使用博客，您的验证码是</h2><h3 style='color:blue'>"+code+"</h3>"
	//将验证码存储到redis中
	model.Redis.Do("set",to,code)
	status:= model.SendEmail(to,"验证码",content)
	c.JSON(http.StatusOK,gin.H{
		"status":status,
		"message":errmsg.GetErrMsg(status),
	})
}

//todo 使用邮件登录
func LoginByEmail(c *gin.Context)  {
	var token string
	var code int
	var user model.User
	to:=c.Query("email")
	userCode:=c.Query("code")
	//从redis中拿到正确的验证码做比较
	serverCode:= model.Redis.Get(to).Val()

	//比较验证码
	if userCode!=serverCode{
		code=errmsg.ERROR_CODE_WRONG
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}
	//通过email查找用户
	user,code= model.GetUserByEmail(to)
	if code == errmsg.SUCCESS {
		//jwt验证
		token, code = middleware.SetToken(user.Username)
	}
	//登录成功删除redis中的email,防止重复使用
	model.Redis.Del(to)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
