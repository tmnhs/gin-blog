package routes

import (
	v1 "demo1/MyBlog/api/v1"
	"demo1/MyBlog/middleware"
	"demo1/MyBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {

	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	//这里是后台管理的页面加载
	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")
	r.StaticFile("admin/favicon.ico", "static/admin/favicon.ico")
	r.GET("admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//这里是博客前台展示的页面加载，如需要可以打开，切记不可同时打开
	//r.LoadHTMLGlob("static/front/index.html")
	//r.Static("front/static", "static/front/static")
	//r.StaticFile("front/favicon.ico", "static/front/favicon.ico")
	//r.GET("front", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})
	//设置中间件，以下操作需要用户权限（凭证）
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{

		//	用户模块的路由接口

		//编辑用户信息
		auth.POST("user/update", v1.EditUser)
		//删除用户
		auth.DELETE("user/:id", v1.DeleteUser)

		//	分类模块的路由接口
		//添加分类
		auth.POST("category/add", v1.AddCategory)

		//编辑分类信息
		auth.PUT("category/:id", v1.EditCategory)
		//删除分类
		auth.DELETE("category/:id", v1.DeleteCategory)

		//	文章模块的路由接口
		//添加文章
		auth.POST("article/add", v1.AddArticle)
		//添加评论
		auth.POST("comment", v1.AddComment)
		//编辑文章
		auth.PUT("article/:id", v1.EditArticle)
		//删除文章
		auth.DELETE("article/:id", v1.DeleteArticle)
		//删除评论
		auth.DELETE("comment/:id", v1.DeleteComment)
		//上传文件
		auth.POST("upload", v1.Upload)

		//更新个人设置
		auth.PUT("profile", v1.UpdateProfile)
	}

	routeV1 := r.Group("api/v1")
	{
		//	添加用户
		routeV1.POST("user/add", v1.AddUser)
		//查询所有用户
		routeV1.POST("users", v1.GetUsers)
		//查询用户详细信息，包括文章
		routeV1.GET("user/:id", v1.GetUserInfo)
		//通过id查询分类信息
		routeV1.GET("category/:id", v1.FindCategoryById)
		//查询所有分类
		routeV1.GET("category", v1.GetCategory)
		//查询所有文章信息
		routeV1.POST("articles", v1.GetArticle)
		//查询某分类下的所有文章
		routeV1.GET("article/cate/:id", v1.GetCateArticle)
		//查询谋篇文章的基本信息
		routeV1.GET("article/info/:id", v1.GetArticleInfo)
		//查询某文章下的所有评论
		routeV1.GET("comment/:id", v1.GetComments)
		//登录
		routeV1.POST("login", v1.Login)
		//注册用户
		routeV1.POST("/register", v1.Register)
		//邮件激活
		routeV1.GET("/active", v1.ActiveEmail)
		//登录发送邮件,需要参数email
		routeV1.GET("sendemail", v1.SendEmailForCode)
		//使用邮箱登录,需要参数email和验证码
		routeV1.GET("loginbyemail", v1.LoginByEmail)

		//获取个人信息
		routeV1.GET("profile", v1.GetProfile)
	}

	r.Run(utils.HttpPort)
}
