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
)

type LoginResponse struct {
	Token string `json:"token"`
}

// @Summary 登录
// @description 用户登录接口
// @tags user
// @Produce json
// @Param userName body string true "用户名"
// @Param passWord body string true "密码"
// @Success 200 {object} app.Response{data=api.LoginResponse} "Success Response"
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /login [post]
func Login(c *gin.Context) {
	appG := app.Gin{C: c}

	// 绑定form，并校验参数
	var form validate.LoginParam
	httpCode, code, errMsg := app.BindAndValid(c, &form)
	if httpCode != http.StatusOK {
		appG.Response(httpCode, code, errMsg)
		return
	}

	// 获取用户信息
	us := userService.User{Username: form.Username}
	user, err := us.GetUserByUsername()
	if err != nil {
		logging.Error(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	// check：用户输入的密码和数据库中的密码
	pass := util.CheckPasswordHash(form.Password, user.Password)
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
// @Param username body string true "用户名"
// @Param passWord body string true "密码"
// @Param rePassword body string true "确认密码"
// @Success 200 {object} app.Response{data=boolean} "Success Response"
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /signup [post]
func SignUp(c *gin.Context) {
	appG := app.Gin{C: c}

	var err error
	var form validate.SignUpParam

	// 校验参数
	// 这个方法不是 线程安全的，需要在每个校验前进行注册
	if err = validate.V.RegisterValidation("SignUpParamUsernameValidate", validate.SignUpParamUsernameValidation); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	httpCode, code, errMsg := app.BindAndValid(c, &form)
	if httpCode != http.StatusOK {
		appG.Response(httpCode, code, errMsg)
		return
	}

	us := userService.User{Username: form.Username, Password: form.Password}
	err = us.Add()
	if err != nil {
		logging.Error(err.Error())
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, true)
}
