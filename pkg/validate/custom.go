package validate

import (
	"github.com/go-playground/validator/v10"
	"ticket-crawler/models"
)

func SignUpParamUsernameValidation(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	exist, _ := models.ExistUserByUsername(username)
	// 如果存在，校验不通过，返回false
	return !exist
}
