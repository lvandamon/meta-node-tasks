package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 记录请求开始时间
		start := time.Now()
		// 继续处理请求
		ctx.Next()
		// 记录请求结束时间
		end := time.Now()
		latency := end.Sub(start)
		// 获取请求方法和路径
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		// 获取响应状态码
		statusCode := ctx.Writer.Status()
		// 记录日志
		logrus.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency":     latency,
			"method":      method,
			"path":        path,
		}).Info("Request processed")
	}
}
