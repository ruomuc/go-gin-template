package main

import (
	"fmt"
	"net/http"
	"ticket-crawler/pkg/setting"
	"ticket-crawler/pkg/validate"
	"ticket-crawler/routers"
)

// @title ticket-crawler API
// @version  1.0
// @contact.name ruomu
// @contact.url blog.seeln.com
// @contact.email 252615299@qq.com
func main() {
	setting.SetUp()
	validate.InitTrans()

	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe()
}
