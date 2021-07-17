package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket-crawler/pkg/e"
	"ticket-crawler/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.GetHeader("authorization")
		if token == "" {
			code = e.AuthorizationnNotFound
		} else {
			customData, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.AuthTokenExpired
				default:
					code = e.AuthTokenFailed
				}
			} else {
				// 把我们需要的信息，放入上下文
				c.Set("extras", customData)
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
