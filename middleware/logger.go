package middleware

import (
	"go-gin-template/pkg/logging"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	logger := logging.Logger
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		// 执行时间
		duration := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method
		// 路由
		reqURI := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Infof("| %d | %v | %15s | %s | %s |",
			statusCode,
			duration,
			clientIP,
			reqMethod,
			reqURI,
		)
	}
}
