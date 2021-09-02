package main

import (
	"fmt"
	"go-gin-template/models"
	"go-gin-template/pkg/logging"
	"go-gin-template/pkg/setting"
	"go-gin-template/pkg/validate"
	"go-gin-template/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title go-gin-template API
// @version  1.0
// @contact.name ruomu
// @contact.url blog.seeln.com
// @contact.email 252615299@qq.com
// @host localhost:8080
// @BasePath /api/v1

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routers := router.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        routers,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe()
}

func init() {
	setting.SetUp()
	logging.Setup()
	models.SetUp()
	validate.InitTrans()
}
