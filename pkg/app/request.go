package app

import (
	"go-gin-template/pkg/logging"
	"go-gin-template/pkg/validate"

	"github.com/go-playground/validator/v10"
)

func MarkErrors(errs validator.ValidationErrors) string {
	newErr := validate.Translate(errs)
	for _, ne := range newErr {
		logging.Logger.Error(ne)
	}
	// 返回一个校验错误，用于返回给前端
	return newErr[0]
}
