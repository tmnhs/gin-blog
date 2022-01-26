package middleware

//跨域参数配置
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	//return cors.New(
	//	cors.Config{
	//		AllowAllOrigins:  true,
	//		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONs"},
	//		AllowHeaders:     []string{"*"},
	//		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
	//		AllowCredentials: true,
	//		MaxAge:           12 * time.Hour,
	//	})
	return cors.Default()

}
