package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"ticket-crawler/pkg/e"
	"ticket-crawler/pkg/validate"
)

// BindAndValid bind and valid form data
func BindAndValid(c *gin.Context, form interface{}) (int, int, interface{}) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.InvalidParam, nil
	}

	// 参数校验
	err = validate.V.Struct(form)
	if err != nil {
		errMsg := MarkErrors(err.(validator.ValidationErrors))
		return http.StatusBadRequest, e.InvalidParam, errMsg
	}
	return http.StatusOK, e.SUCCESS, "ok"
}
