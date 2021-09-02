package logging

import (
	"fmt"
	"go-gin-template/pkg/file"
	"go-gin-template/pkg/setting"
	"io"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type Level uint32

var (
	f      *os.File
	Logger *logrus.Logger
)

// Setup initialize the log instance
func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	f, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}
	// gin的框架的路由日志
	//gin.DisableConsoleColor()
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// logrus日志对象初始化
	Logger = logrus.New()
	// 同时将日志输出的文件和标准错误流中
	Logger.SetOutput(io.MultiWriter(os.Stdout, f))
	// 设置在输出日志中添加文件名和方法信息
	Logger.SetReportCaller(true)
	// 日志格式
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.AppSetting.LogSaveName,
		time.Now().Format(setting.AppSetting.TimeFormat),
		setting.AppSetting.LogFileExt,
	)
}
