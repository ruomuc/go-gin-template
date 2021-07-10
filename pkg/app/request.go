package app

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"ticket-crawler/pkg/validate"
)

func MakeErrors(errs validator.ValidationErrors) string {
	newErr := validate.Translate(errs)
	for _, ne := range newErr {
		fmt.Println(ne)
	}
	// 返回一个校验错误，用于返回给前端
	return newErr[0]
}
