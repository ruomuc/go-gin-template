package models

import "github.com/jinzhu/gorm"

// 结构体名称要和表名相同
// 因为我已经全局禁用了 gorm 的复数表名
type Users struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
}

func ExistUserByUsername(username string) (bool, error) {
	var user Users
	err := db.Select("id").Where("username = ? AND is_deleted = 0", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func AddUser(data map[string]interface{}) error {
	user := Users{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Phone:    data["phone"].(int),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
