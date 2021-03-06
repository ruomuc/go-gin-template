package router

import (
	"go-gin-template/middleware"
	"go-gin-template/router/api"
	v1 "go-gin-template/router/api/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 加载中间件
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	// 通用路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/signup", api.SignUp)
	r.POST("/login", api.Login)

	// 路由组
	apiv1 := r.Group("/api/v1")
	// 鉴权中间件加载到路由组上，因为通用路由不用鉴权。。
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/test-token", v1.TestToken)
	}
	return r
}
