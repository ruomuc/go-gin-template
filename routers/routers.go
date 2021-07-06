package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket-crawler/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 通用路由
	r.GET("/signin", api.SignIn)
	r.GET("/login", api.Login)

	// 路由组
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, "helloWorld")
		})
	}
	return r
}
