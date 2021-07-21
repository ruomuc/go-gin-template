package v1

import (
	"net/http"
	"ticket-crawler/pkg/app"
	"ticket-crawler/pkg/e"
	"ticket-crawler/pkg/validate"
	ticket_service "ticket-crawler/service/ticket-service"

	"github.com/gin-gonic/gin"
)

// @Summary 获取机票价格源数据
// @description 通过这个接口，去第三方获取机票价格
// @tags ticket
// @Produce json
// @Param fromCity body string true "出发地"
// @Param toCity body string true "目的地"
// @Param startDate body string true "开始日期"
// @Param endDate body string true "截止日期"
// @Success 200 {object} app.Response{data=boolean} "Success Response"
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /ticket [post]
func ClimbTicketInfo(c *gin.Context) {
	appG := app.Gin{C: c}

	// 表单绑定&参数校验
	var form validate.ClimbTicketInfoParam
	httpCode, code, errMsg := app.BindAndValid(c, &form)
	if httpCode != http.StatusOK {
		appG.Response(httpCode, code, errMsg)
		return
	}

	ts := ticket_service.TicketRequest{
		FromCity:  form.FromCity,
		ToCity:    form.ToCity,
		StartDate: form.StartDate,
		EndDate:   form.EndDate,
	}
	ts.Climb()
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
