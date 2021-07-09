package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	_ "ticket-crawler/docs"
	"ticket-crawler/pkg/app"
	"ticket-crawler/pkg/e"
	"ticket-crawler/pkg/validate"
	userService "ticket-crawler/service/user-service"
)

type LoginResponse struct {
	Token string `json:"token"`
}

// @Summary 登录
// @description 用户登录接口
// @tags user
// @Produce json
// @Param userName query string true "用户名"
// @Param passWord query string true "密码"
// @Success 200 {object} app.Response{data=api.LoginResponse} "desc"
// @Router /login [post]
func Login(c *gin.Context) {

}

type User struct {
	Username string `validate:"required,min=1,max=20"`
	Password string `validate:"required,min=8,max=64"`
}

func UsernameDuplicate(v *validator.Validate) bool {
	return true
}

// @Summary 注册
// @description 用户注册接口
// @tags user
// @Produce json
// @Param username query string true "用户名"
// @Param passWord query string true "密码"
// @Success 200 {object} app.Response{data=boolean} "desc"
// @Failure 500 {object} app.Response
// @Router /signin [post]
func SignIn(c *gin.Context) {
	appG := app.Gin{C: c}

	var err error

	username := c.Query("username")
	password := c.Query("password")

	// 校验参数
	u := &User{Username: username, Password: password}
	err = validate.Validator.Struct(u)
	if err != nil {
		newErr := validate.Translate(err.(validator.ValidationErrors))
		app.MakeErrors(newErr)
		appG.Response(http.StatusBadRequest, e.InvalidParam, newErr[0])
		return
	}

	us := userService.User{Username: username, Password: password}
	err = us.AddUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, true)
}
