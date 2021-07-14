package api

import (
	"net/http"
	_ "ticket-crawler/docs"
	"ticket-crawler/pkg/app"
	"ticket-crawler/pkg/e"
	"ticket-crawler/pkg/logging"
	"ticket-crawler/pkg/util"
	"ticket-crawler/pkg/validate"
	userService "ticket-crawler/service/user-service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	appG := app.Gin{C: c}

	username := c.Query("username")
	password := c.Query("password")
	// 参数校验
	loginParam := &validate.LoginParam{Username: username, Password: password}
	err := validate.V.Struct(loginParam)
	if err != nil {
		errMsg := app.MarkErrors(err.(validator.ValidationErrors))
		appG.Response(http.StatusInternalServerError, e.ERROR, errMsg)
		return
	}

	// 获取用户信息
	us := userService.User{Username: username}
	user, err := us.GetUserByUsername()
	if err != nil {
		logging.Error(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	// check：用户输入的密码和数据库中的密码
	pass := util.CheckPasswordHash(password, user.Password)
	if !pass {
		appG.Response(http.StatusUnauthorized, e.ErrorPassword, nil)
		return
	}
	// 生成jwt token
	token, err := util.GenerateToken(user.ID, user.Username, user.Phone)
	if err != nil {
		logging.Error(err.Error())
		appG.Response(http.StatusInternalServerError, e.TokenGenerateFailed, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, LoginResponse{Token: token})
}

// @Summary 注册
// @description 用户注册接口
// @tags user
// @Produce json
// @Param username query string true "用户名"
// @Param passWord query string true "密码"
// @Success 200 {object} app.Response{data=boolean} "desc"
// @Failure 500 {object} app.Response
// @Router /signup [post]
func SignUp(c *gin.Context) {
	appG := app.Gin{C: c}

	var err error

	username := c.Query("username")
	password := c.Query("password")
	rePassword := c.Query("rePassword")

	// 校验参数
	// 这个方法不是 线程安全的，需要在每个校验前进行注册
	if err = validate.V.RegisterValidation("SignUpParamUsernameValidate", validate.SignUpParamUsernameValidation); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	signUpParam := &validate.SignUpParam{Username: username, Password: password, RePassword: rePassword}
	err = validate.V.Struct(signUpParam)
	if err != nil {
		msg := app.MarkErrors(err.(validator.ValidationErrors))
		appG.Response(http.StatusBadRequest, e.InvalidParam, msg)
		return
	}

	us := userService.User{Username: username, Password: password}
	err = us.Add()
	if err != nil {
		logging.Error(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, true)
}
