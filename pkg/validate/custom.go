package validate

import (
	"github.com/go-playground/validator/v10"
	"ticket-crawler/models"
)

func SignUpParamUsernameValidation(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	exist, _ := models.ExistUserByUsername(username)
	return exist
}
