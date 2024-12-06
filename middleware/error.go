package middleware

import (
	"bi-activity/response"
	serror "bi-activity/response/errors"
	"errors"
	"github.com/gin-gonic/gin"
)

// errorHandler 统一错误处理中间件 [可选择使用]
func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, e := range c.Errors {
			err := e.Err
			var selfErr serror.SelfError
			if errors.As(err, &selfErr) {
				status, resp := response.Failf(selfErr, "参数错误")
				c.JSON(status, resp)
				return
			}

			c.JSON(500, gin.H{
				"error": err,
				"msg":   "服务器内部错误",
			})
		}
	}
}
