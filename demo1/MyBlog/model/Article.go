package model

//文章模型

import (
	"demo1/MyBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img     string `gorm:"type:varchar(200)" json:"img"`
	Category Category `gorm:"foreignkey:Cid"`
	Comments []Comment
}



//todo 添加文章,此时需要建立中间表
func CreateArticle(data *Article) int {

	//密码加密
	//data.Password=ScryptPw(data.Password)
	//添加中间表

	err = db.Create(&data).Error
	//_=db.Create(&userArticle).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//todo 添加评论
func CreateComment(id int,data Comment)int {
	var article Article
	err=db.Model(&Article{}).Where("id = ?",id).First(&article).Error
	if err != nil {
		return errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	article.Comments = append(article.Comments, data)
	err = db.Save(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//todo 查询文章下的所有评论
func GetCommentsByArticleId(id int)([]Comment,int)  {
	var comments []Comment
	var article Article
	var _ = db.Model(&Article{}).Where("id = ?",id).First(&article).Error
	err = db.Where("article_id = ?", id).Find(&comments).Error
	if article.ID==0{
		return nil ,errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	if len(comments)==0 {
		return nil,errmsg.ERROR_ARTICLE_NO_COMMENTS
	}
	return comments,errmsg.SUCCESS
}

//todo 查询文章列表，带分页效果
func GetArticles(pageSize int,pageNum int )([]Article,int,int) {
	var articles []Article
	var total int
	err=db.Debug().Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&articles).Error
	//查找总数
	_=db.Debug().Find(&[]Article{}).Count(&total).Error
	if err != nil &&err!=gorm.ErrRecordNotFound{
		return nil,errmsg.ERROR,0
	}
	return articles,errmsg.SUCCESS,total
}

// todo 查询分类下的所有文章
func GetCateArticle(id int,pageSize int,pageNum int )([]Article,int,int)  {
	var cateArticleList []Article
	var total int
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArticleList).Error
	//查找总数
	_ = db.Limit(pageSize).Offset(-1).Where("cid = ?", id).Find(&[]Article{}).Count(&total).Error
	if err != nil {
		return cateArticleList,errmsg.ERROR_CATE_NOT_EXIST,0
	}
	return cateArticleList,errmsg.SUCCESS,total
}

//todo 查询单个文章详细信息
func GetArticleInfo(id int)(Article,int){
	var article Article
	//根据文章id查找出其评论
	comments,_:= GetCommentsByArticleId(id)

	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	for _,comment:=range comments{
		article.Comments=append(article.Comments,comment)
	}
	if err != nil {
		return article,errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article,errmsg.SUCCESS
}
//todo 编辑文章信息
func EditArticle(id int,data *Article) int  {
	//利用map集合更新
	var maps=make(map[string]interface{})
	maps["title"]=data.Title
	maps["cid"]=data.Cid
	maps["desc"]=data.Desc
	maps["content"]=data.Content
	maps["img"]=data.Img
	err = db.Model(&Article{}).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
//todo 删除文章
func DeleteArticle( id int)  int{
	//删除文章时，需要把与之相关联的中间表也删除
	DeleteMidByArticleId(id)
	err= db.Where("id = ?",id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
