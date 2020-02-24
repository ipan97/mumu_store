package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ipan97/mumu-store/app"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Initialize(r *gin.Engine, db *gorm.DB) {
	// Web Route
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Mumu Store",
		})
	})

	// Instance dependency injection
	di := app.NewDependencyInjection(db)

	// API Route
	api := r.Group("/api")
	{
		categoryController := di.InjectCategoryController()
		api.GET("/categories", categoryController.FindAllCategories)
	}

}
