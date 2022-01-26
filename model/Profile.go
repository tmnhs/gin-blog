package model

import "demo1/MyBlog/utils/errmsg"

type Profile struct {
	ID int  `gorm:"primaryKey" json:"id"`
	Name string  `gorm:"type:varchar(20)" json:"name"`
	Desc string  `gorm:"type:varchar(200)" json:"desc"`
	QqChat string  `gorm:"type:varchar(32)" json:"qq_chat"`
	WeChat string  `gorm:"ype:varchar(32)" json:"wechat"`
	Weibo string  `gorm:"ype:varchar(32)" json:"weibo"`
	Email string  `gorm:"ype:varchar(32)" json:"email"`
	Img string  `gorm:"ype:varchar(80)" json:"img"`
	Avatar string  `gorm:"ype:varchar(80)" json:"avatar"`

}

//获取个人信息
func GetProfile(id int) (*Profile,int) {
	var profile  Profile
	err=db.Model(&profile).Where("ID = ?",id).First(&profile).Error
	if err != nil {
		return nil,errmsg.ERROR
	}
	return &profile,errmsg.SUCCESS
}

//更新个人信息
func UpdateProfile(id int ,profile * Profile)  int{
	err=db.Model(&Profile{}).Where("ID = ?",id).Update(profile).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}