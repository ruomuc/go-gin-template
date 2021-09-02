package validate

import (
	"go-gin-template/models"

	"github.com/go-playground/validator/v10"
)

func SignUpParamUsernameValidation(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	exist, _ := models.ExistUserByUsername(username)
	// 如果存在，校验不通过，返回false
	return !exist
}
