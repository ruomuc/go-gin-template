package models

import "github.com/jinzhu/gorm"

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
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}
