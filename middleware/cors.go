package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func Cors() gin.HandlerFunc {
//	//return cors.New(
//	//	cors.Config{
//	//		AllowAllOrigins:  true,
//	//		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONs"},
//	//		AllowHeaders:     []string{"*"},
//	//		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
//	//		AllowCredentials: true,
//	//		MaxAge:           12 * time.Hour,
//	//	})
//	return cors.Default()
//
//}
//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Header("Access-Control-Allow-Origin", "true")
//		c.Header("Access-Control-Allow-Headers", "*")
//		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS")
//		c.Header("Access-Control-Allow-Credentials", "true")
//		c.Header("Access-Control-Expose-Headers", "*")
//		if c.Request.Method == "OPTIONS" {
//			c.JSON(http.StatusOK, "")
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//
//}

// Cors 直接放行所有跨域请求并放行所有 OPTIONS 方法
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Content-Length, Authorization, Token,X-Token,X-User-Id,x-requested-with")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//Access to XMLHttpRequest at 'http://localhost/api/v1/upload' from origin 'http://localhost:8080' has been blocked by CORS policy: Request header field x-requested-with is not allowed by Access-Control-Allow-Headers in preflight response.
		//AddArt.vue?e24c:105
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
