package v1

import (
	"demo1/MyBlog/model"
	"demo1/MyBlog/proto"
	"demo1/MyBlog/utils/errmsg"
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

	userArticle.ArticleId = data.ID

	userName, _ := c.Get("username")
	userArticle.UserId = model.FindUserIdByName(userName)
	_ = model.CreateUserArticle(&userArticle)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// todo 添加评论
func AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data)
	id := data.ArticleID

	code := model.CreateComment(id, data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 删除评论
func DeleteComment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := model.DeleteComment(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 查询单个文章
func GetArticleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetArticleInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 查询文章下的所有评论
func GetComments(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCommentsByArticleId(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo 查询某分类下的所有文章
func GetCateArticle(c *gin.Context) {
	//获取参数
	//var req proto.ReqCommon
	id, _ := strconv.Atoi(c.Param("id"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))

	//_=c.ShouldBindJSON(&req)
	if pageNum == 0 {
		pageNum = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	//fmt.Println(req)
	data, code, total := model.GetCateArticle(id, pageSize, pageNum)
	articles := make([]*proto.RspFindArticle, 0)
	for i := 0; i < len(data); i++ {
		article := &proto.RspFindArticle{
			ID:        data[i].ID,
			Title:     data[i].Title,
			CreatedAt: data[i].CreatedAt,
			Desc:      data[i].Desc,
			Content:   data[i].Content,
			Img:       data[i].Img,
		}
		Cate, err := model.FindCategoryById(data[i].Cid)
		if err != nil {
			article.Name = ""
		} else {
			article.Name = Cate.Name
		}
		articles = append(articles, article)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//todo  查询文章列表
func GetArticle(c *gin.Context) {
	//获取参数
	var req proto.ReqFindArticle
	_ = c.ShouldBindJSON(&req)
	//如果pageSize等于0，则取消分页
	if req.PageSize == 0 {
		req.PageSize = -1
	}
	if req.PageNum == 0 {
		req.PageNum = -1
	}
	//调用函数查看所有用户
	data, code, total := model.GetArticles(req.Title, req.PageSize, req.PageNum)

	articles := make([]*proto.RspFindArticle, 0)
	for i := 0; i < len(data); i++ {
		article := &proto.RspFindArticle{
			ID:        data[i].ID,
			Title:     data[i].Title,
			CreatedAt: data[i].CreatedAt,
			Desc:      data[i].Desc,
			Content:   data[i].Content,
			Img:       data[i].Img,
		}
		Cate, err := model.FindCategoryById(data[i].Cid)
		if err != nil {
			article.Name = ""
		} else {
			article.Name = Cate.Name
		}
		articles = append(articles, article)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"total":   total,
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
