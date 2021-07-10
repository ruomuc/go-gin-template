package validate

import (
	"github.com/go-playground/validator/v10"
)

func SignUpParamPasswordValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(SignUpParam)
	if user.Password != user.RePassword {
		sl.ReportError(user.RePassword, "Password", "password", "password", "")
	}
}
