package model

import (
	"github.com/gin-gonic/gin"
)

const SessionName = "session"

//设置客户端session_id
func SetSession(c *gin.Context, sessionID string, expireTime int) {
	c.SetCookie(SessionName, sessionID, expireTime, "/", "localhost", false, true)
	//http.SetCookie(c.Writer, &http.Cookie{
	//	Name:     SessionName,
	//	Value:    sessionID,
	//	MaxAge:   expireTime,
	//	Path:     "/",
	//})
}

func ClearSession(c *gin.Context) {
	c.SetCookie(SessionName, "", -1, "/", "localhost", false, true)
}
