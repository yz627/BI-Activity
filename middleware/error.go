package middleware

import (
	"bi-activity/reponse"
	serror "bi-activity/reponse/errors"
	"errors"
	"github.com/gin-gonic/gin"
)

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, e := range c.Errors {
			err := e.Err
			var selfErr serror.SelfError
			if errors.As(err, &selfErr) {
				status, resp := reponse.Failf(selfErr, "参数错误")
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
