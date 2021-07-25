package user_service

import (
	"ticket-crawler/models"
	"ticket-crawler/pkg/util"
)

type User struct {
	Username string
	Password string
	Phone    int
}

// 添加一个用户
func (u *User) Add() error {
	// 密码需要加密
	password, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}
	user := map[string]interface{}{
		"username": u.Username,
		"password": password,
		"phone":    u.Phone,
	}
	return models.AddUser(user)
}

func (u *User) GetUserByUsername() (*models.User, error) {
	return models.GetUserByUsername(u.Username)
}
