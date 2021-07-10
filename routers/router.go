package router

import (
	"net/http"
	"ticket-crawler/routers/api"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 通用路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/signup", api.SignUp)
	r.POST("/login", api.Login)

	// 路由组
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, "helloWorld")
		})
	}
	return r
}
