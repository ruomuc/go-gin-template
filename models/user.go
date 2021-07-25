package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// 结构体名称要和表名相同
// 因为我已经全局禁用了 gorm 的复数表名
type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}

func ExistUserByUsername(username string) (bool, error) {
	var user User
	err := db.Select("id").Where("username = ? AND is_deleted = 0", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, errors.Wrap(err, "ExistUserByUsername->db.Select error")
	}
	return user.ID > 0, nil
}

func AddUser(data map[string]interface{}) error {
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Phone:    data["phone"].(int),
	}
	err := db.Create(&user).Error
	return errors.Wrap(err, "AddUser->db.Create error")
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.Where("username = ? and is_deleted = 0", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.Wrap(err, "GetUserByUsername error")
	}
	return &user, nil
}
