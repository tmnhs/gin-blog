package model

//评论模型

import (
	"demo1/MyBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Commentator string `gorm:"type:varchar(20);not null" json:"commentator"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	ArticleID int `gorm:"type:int;not null" json:"article_id"`
	ParentID int `gorm:"type:int " json:"parent_id"`
}

//todo 根据id删除评论
func DeleteComment(id int) int {
	err=db.Where("id = ?",id).Delete(&Comment{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}