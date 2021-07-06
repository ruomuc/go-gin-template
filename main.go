package main

import (
	"fmt"
	"net/http"
	"ticket-crawler/pkg/setting"
	"ticket-crawler/routers"
)

func main() {
	setting.SetUp()
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
