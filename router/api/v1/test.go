package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"ticket-crawler/pkg/app"
	"ticket-crawler/pkg/e"
	"ticket-crawler/pkg/logging"
)

// TestToken test Authorization param in header
func TestToken(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		data interface{}
		ok   bool
	)

	if data, ok = c.Get("extras"); ok {
		logging.Info(fmt.Sprintf("%+v", data))
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
