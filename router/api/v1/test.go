package v1

import (
	"fmt"
	"go-gin-template/pkg/app"
	"go-gin-template/pkg/e"
	"go-gin-template/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestToken test Authorization param in header
func TestToken(c *gin.Context) {
	appG := app.Gin{C: c}
	var (
		data interface{}
		ok   bool
	)

	if data, ok = c.Get("extras"); ok {
		logging.Logger.Info(fmt.Sprintf("%+v", data))
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
