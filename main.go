package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket-crawler/models"
	"ticket-crawler/pkg/logging"
	"ticket-crawler/pkg/setting"
	"ticket-crawler/pkg/validate"
	"ticket-crawler/router"
)

// @title ticket-crawler API
// @version  1.0
// @contact.name ruomu
// @contact.url blog.seeln.com
// @contact.email 252615299@qq.com
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
