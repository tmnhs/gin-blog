package model

import (
	"demo1/MyBlog/utils"
	"demo1/MyBlog/utils/errmsg"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(20)" json:"name"`
	Desc   string `gorm:"type:varchar(200)" json:"desc"`
	QqChat string `gorm:"type:varchar(32)" json:"qq_chat"`
	WeChat string `gorm:"ype:varchar(32)" json:"wechat"`
	Weibo  string `gorm:"type:varchar(32)" json:"weibo"`
	Email  string `gorm:"type:varchar(32)" json:"email"`
	Img    string `gorm:"type:varchar(80)" json:"img"`
	Avatar string `gorm:"type:varchar(80)" json:"avatar"`
}

//获取个人信息
func GetProfileById(id int) (*Profile, int) {
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).First(&profile).Error
	//profile:=Redis.Get("profile")
	if err != nil {
		return nil, errmsg.ERROR
	}
	return &profile, errmsg.SUCCESS
}

//获取个人信息
func GetProfile(c *gin.Context) (*Profile, int) {
	var profile *Profile
	sessionId, err := c.Cookie(SessionName)
	//fmt.Println("seesionId:",sessionId)
	if err != nil {
		return nil, errmsg.ERROR
	}
	result, err := Redis.Get(sessionId).Result()
	if err != nil {
		return nil, errmsg.ERROR
	}
	err = json.Unmarshal([]byte(result), &profile)
	if err != nil {
		fmt.Println("unmarshall json false")
	}
	return profile, errmsg.SUCCESS
}

//更新个人信息
func UpdateProfile(c *gin.Context, id int, profile *Profile) int {
	var _profile *Profile
	err = db.Model(&Profile{}).Where("ID = ?", id).Update(profile).Error
	if err != nil {
		return errmsg.ERROR
	}
	sessionId, err := c.Cookie(SessionName)
	if err != nil {
		return errmsg.ERROR
	}
	result, err := Redis.Get(sessionId).Result()
	if err != nil {
		return errmsg.ERROR
	}
	err = json.Unmarshal([]byte(result), &_profile)
	if err != nil {
		fmt.Println("unmarshall json false")
	}
	if _profile.ID == profile.ID {
		//需要更新redis
		profileJson, _ := json.Marshal(profile)
		_, err = Redis.Set(sessionId, string(profileJson), utils.Expiration).Result()
		if err != nil {
			return errmsg.ERROR
		}
	}
	return errmsg.SUCCESS
}

func CreateProfile(profile *Profile) int {
	err = db.Model(&Profile{}).Create(&profile).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
