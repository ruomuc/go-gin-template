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
	if err := models.AddUser(user); err != nil {
		return err
	}
	return nil
}
