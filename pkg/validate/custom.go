package validate

import (
	"github.com/go-playground/validator/v10"
)

func SignUpParamUsernameValidation(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	if username == "zm" {
		return false
	}
	return true
}
