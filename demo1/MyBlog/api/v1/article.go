package v1

import (
	"demo1/MyBlog/model"
	"demo1/MyBlog/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
文章处理模块的接口
*/

//todo 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	var userArticle model.UserArticle
	//绑定
	_ = c.ShouldBind(&data)
	code := model.CreateArticle(&data)

	userArticle.ArticleId= data.ID

	userName, _ :=c.Get("username")
	//fmt.Printf("%T\n",username)
	//fmt.Println(username)
	userArticle.UserId =model.FindUserIdByName(userName)
	fmt.Println(userArticle)
	_= model.CreateUserArticle(&userArticle)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
// todo 添加评论
func AddComment(c * gin.Context)  {
	var data model.Comment
	_=c.ShouldBindJSON(&data)
	id:=data.ArticleID

	code:=model.CreateComment(id,data)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
//todo 删除评论
func DeleteComment(c * gin.Context){
	id,_:=strconv.Atoi(c.Param("id"))

	code:=model.DeleteComment(id)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//todo 查询单个文章
func GetArticleInfo(c *gin.Context) {
	id,_:=strconv.Atoi(c.Param("id"))

	data,code:=model.GetArticleInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
//todo 查询文章下的所有评论
func GetComments(c *gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))

	data, code:=model.GetCommentsByArticleId(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}
//todo 查询某分类下的所有文章
func GetCateArticle(c *gin.Context)  {
	//获取参数
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	cid,_:=strconv.Atoi(c.Param("cid"))
	//如果pageSize等于0，则取消分页
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data,code,total:=model.GetCateArticle(cid,pageSize,pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":total,
		"message": errmsg.GetErrMsg(code),
	})
}
//todo  查询文章列表
func GetArticle(c *gin.Context) {
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
	data ,code,total:= model.GetArticles(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":total,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 编辑文章
func EditArticle(c *gin.Context) {
	var data model.Article
	//获取参数
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 删除文章
func DeleteArticle(c *gin.Context) {
	//获取参数
	id, _ := strconv.Atoi(c.Param("id"))

	//调用删除的函数
	code := model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
