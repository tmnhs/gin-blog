package middleware

//跨域参数配置
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"Origin"},
			ExposeHeaders: []string{"Content-Length","Authorization"},
			MaxAge: 12*time.Hour,
		})
	}
}