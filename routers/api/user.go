package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "ticket-crawler/docs"
	"ticket-crawler/pkg/app"
	"ticket-crawler/pkg/e"
	"ticket-crawler/pkg/util"
	user_service "ticket-crawler/service/user-service"
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

type user struct {
	username string `validate:"max=20"`
	password string `validate:"max=64"`
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
	//valid := validator.New()
	username := c.Query("username")
	password := c.Query("password")

	// 校验参数
	//u := user{username: username, password: password}

	// 加密密码
	password, err = util.HashPassword(password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	userService := user_service.User{Username: username, Password: password}
	// 判断用户名是否存在
	//exist, err = userService.ExistUserByUsername()

	err = userService.AddUser()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, true)
}
