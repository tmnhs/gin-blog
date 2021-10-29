package model

import (
	"demo1/MyBlog/utils/errmsg"
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Email string `gorm:"type:varchar(20);not null" json:"email" validate:"required,email" label:"邮箱"`
	Role     int    `gorm:"type:int DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"`
	Article []Article `gorm:"many2many:user_article"`
	Status string `gorm:"DEFAULT:'N'"`		//激活状态
	Code string `gorm:"type:varchar(80)"`	//激活码
}

//todo 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USERD //1001
	}
	return errmsg.SUCCESS
}

//todo 添加用户
func CreateUser(data *User) int {
	//密码加密
	//data.Password=ScryptPw(data.Password)
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS
}

//todo 查询用户详细信息，包括文章
func GetUserInfo(id int) (User,int) {
	var user User

	err=db.Debug().Model(&user).Preload("Article").Where("id = ?",id).First(&user).Error
	if err != nil {
		return  user,errmsg.ERROR
	}
	return user,errmsg.SUCCESS
}

//todo 查询用户列表，带分页效果
func GetUsers(pageSize int,pageNum int )([]User,int) {
	var users []User
	var total int
	err=db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Count(&total).Error
	if err != nil &&err!=gorm.ErrRecordNotFound{
		return nil,0
	}
	return users,total
}

//todo 编辑用户
func EditUser(id int,data *User) int  {
	var maps=make(map[string]interface{})
	maps["username"]=data.Username
	maps["role"]=data.Role
	maps["email"]=data.Email
	err = db.Model(&User{}).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
//todo 删除用户
func DeleteUser( id int)  int{
	//删除与该用户相关联的中间表
	DeleteMidByUserId(id)

	err= db.Where("id = ?",id).Delete(&User{}).Error

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
//在调用函数之前执行
func (u *User) BeforeSave()  {
	u.Password=ScryptPw(u.Password)
}
// todo 使用scrypt密码加密
func ScryptPw(password string)string  {
	const KeyLen = 10
	salt:=make([]byte,8)
	salt=[]byte{12,32,14,6,66,22,43,11}
	//加密
	HashPw,err:= scrypt.Key([]byte(password),salt,16384,8,1,KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	//将加密后的密码转化为字符串
	fpw:=base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

//todo 登录验证
func CheckLogin(username string,password string) int {
	var user User
	db.Debug().Where("username = ?",username).First(&user)
	//fmt.Println(user)
	if user.Status=="N"{
		return errmsg.ERROR_EMAIL_HAVE_NOT_ACTIVE
	}
	//判断该用户是否存在
	if user.ID==0{
		return errmsg.ERROR_USER_NOT_EXIST
	}
	//判断用户密码是否正确
	if ScryptPw(password)!=user.Password{
		return errmsg.ERROR_PASSWORD_WRONG
	}
	//用户无权限
	//if user.Role!=2{
	//	return errmsg.ERROR_USER_NO_RIGHT
	//}
	return errmsg.SUCCESS
}

//todo 通过用户名查找用户id
func FindUserIdByName(username interface{}) uint {

	var user User

	db.Where("username = ?", username).First(&user)

	return user.ID

}

//todo 通过激活码查找用户并且激活
func UpdateUserStatus(code string)  int {

	err= db.Debug().Model(&User{}).Where("code = ?",code).UpdateColumn(map[string]interface{}{
		"status":"Y",
	}).Error
	if err != nil {
		return errmsg.ERROR_EMAIL_ACTIVE
	}
	//激活
	return errmsg.SUCCESS
}

//todo 通过邮箱查找用户
func GetUserByEmail(email string) (User, int) {
	var user User

	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCESS
}