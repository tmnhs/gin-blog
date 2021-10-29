package middleware

import (
	"demo1/MyBlog/utils"
	"demo1/MyBlog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwyKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成token
func SetToken(username string) (string, int) {
	//设置到期时间
	expireTime := time.Now().Add(24 * time.Hour) //24小时
	//创建Jwt声明
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	//使用用于签名的算法和令牌
	reqClain := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	//创建jwt字符串
	token, err := reqClain.SignedString(JwyKey)
	//如果出错，则返回服务器内部错误
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//验证token
func CheckToken(token string) (*MyClaims, int) {

	//解析jwt字符串并将结果存储
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwyKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims);  setToken.Valid {
		return key, errmsg.SUCCESS
	}
	return nil, errmsg.ERROR
}

//jwt中间件
var code int

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")

		//如果token为空，返回错误
		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST		//1004
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		//token类型有错误
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//验证token
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//判断时间是否过期
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIEM
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//成功情况下
		c.Set("username", key.Username)
		//c.Set("userid",key.Id)
		c.Next()
	}
}
