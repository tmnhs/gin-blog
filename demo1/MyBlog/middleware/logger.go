package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)
//日志
func Logger() gin.HandlerFunc {
	//设置日志文件的路径
	filePath:="log/log"
	//建立软连接，需要管理员权限
	linkName:="latest_log.log"
	src,err:=os.OpenFile(filePath,os.O_RDWR|os.O_CREATE,0755)
	if err != nil {
		fmt.Println("err: ",err)
	}
	//创建日志
	logger := logrus.New()
	//输出
	logger.Out=src
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//添加时间分割
	logWriter,_:=retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),		//日志保留时间：一周
		retalog.WithRotationTime(24*time.Hour),	//24小时分割一次
		retalog.WithLinkName(linkName),			//建立软连接
		)
	writeMap:=lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	//实例化
	Hook:=lfshook.NewHook(writeMap,&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unkown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode>=500{
			entry.Error()
		}else if statusCode>=400{
			entry.Warn()
		}else{
			entry.Info()
		}
	}
}
