package model

import (
	"demo1/MyBlog/utils/errmsg"
)

type UserArticle struct {
	ArticleId uint `gorm:"type:not null" json:"article_id"`
	UserId uint `gorm:"type:not null" json:"user_id"`
}

//todo 创建中间表
func CreateUserArticle(userArticle *UserArticle) int {
	err=db.Create(&userArticle).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}
//todo 根据用户id删除中间表
func DeleteMidByUserId(id int) int {
	err=db.Where("user_id = ?",id).Delete(&UserArticle{}).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}
//todo 根据文章id删除中间表
func DeleteMidByArticleId(id int) int {
	err=db.Where("article_id = ?",id).Delete(&UserArticle{}).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}
