package app

import (
	"github.com/go-playground/validator/v10"
	"ticket-crawler/pkg/logging"
	"ticket-crawler/pkg/validate"
)

func MarkErrors(errs validator.ValidationErrors) string {
	newErr := validate.Translate(errs)
	for _, ne := range newErr {
		logging.Logger.Error(ne)
	}
	// 返回一个校验错误，用于返回给前端
	return newErr[0]
}
