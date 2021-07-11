package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"ticket-crawler/pkg/setting"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Creator   int       `json:"creator"`
	Reviser   int       `json:"reviser"`
	IsDeleted int       `json:"isDeleted"`
}

func SetUp() {
	var err error
	_, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
	))
	if err != nil {
		log.Fatalf("models setup err: %v", err)
	}

	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(100)
}
