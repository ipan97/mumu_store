package admin

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"

	"github.com/gin-gonic/gin"
)

type DashboardController struct {
}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

func (hc *DashboardController) Index(ctx *gin.Context) {
	ginview.HTML(ctx, http.StatusOK, "index", gin.H{
		"title": "Welcome Admin.",
	})
}
