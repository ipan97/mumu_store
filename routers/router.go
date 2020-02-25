package routers

import (
	"html/template"
	"net/http"

	"github.com/foolin/goview"

	"github.com/foolin/goview/supports/ginview"

	"github.com/gin-gonic/gin"
	"github.com/ipan97/mumu-store/app"
	"github.com/jinzhu/gorm"
)

func Initialize(r *gin.Engine, db *gorm.DB) {
	// Web Route
	r.Static("/public", "./public")
	r.HTMLRender = ginview.New(goview.Config{
		Root:         "./app/views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        make(template.FuncMap),
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Mumu Store",
		})
	})

	// Instance dependency injection
	di := app.NewDependencyInjection(db)

	// Admin Route
	viewAdmin := ginview.NewMiddleware(goview.Config{
		Root:         "./app/views/admin",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        make(template.FuncMap),
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})
	admin := r.Group("/admin", viewAdmin)
	{
		dashboardController := di.InjectDashboardController()
		admin.GET("/", dashboardController.Index)
	}

	// API Route
	api := r.Group("/api")
	{
		// Category
		categoryController := di.InjectCategoryController()
		api.GET("/categories", categoryController.FindAllCategories)

		// Brand
		brandController := di.InjectBrandController()
		api.GET("/brands", brandController.FindAllBrands)
		api.GET("/brands/:id", brandController.FindBrandByID)
	}

}
