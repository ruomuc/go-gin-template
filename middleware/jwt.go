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
		var code int
		token := c.GetHeader("authentication")
		if token == "" {
			code = e.AuthenticationNotFound
		} else {
			err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.AuthTokenExpired
				default:
					code = e.AuthTokenFailed
				}
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
