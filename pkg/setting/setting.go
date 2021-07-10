package setting

import (
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

type App struct {
	PageSize          int
	JwtSecret         string
	ValidatorLanguage string
	RuntimeRootPath   string
	LogSavePath       string
	LogSaveName       string
	TimeFormat        string
	LogFileExt        string
}

type Server struct {
	RunMode string

	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var (
	AppSetting      = &App{}
	ServerSetting   = &Server{}
	DatabaseSetting = &Database{}
)

func SetUp() {
	var cfg *ini.File
	var err error

	profile := os.Getenv("profile")
	switch profile {
	case "prod":
		cfg, err = ini.Load("conf/prod.ini")
	default:
		cfg, err = ini.Load("conf/dev.ini")
	}

	if err != nil {
		log.Fatalf("Fail to load 'conf/dev.ini': %v", err)
	}

	// 通用配置
	err = cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("cfg.Mapto AppSetting err: %v", err)
	}

	// 服务配置
	err = cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("cfg.Mapto ServerSetting err: %v", err)
	}
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

	// 数据库配置
	err = cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("cfg.Mapto DatabaseSetting err: %v", err)
	}
}
