package routers

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/ipan97/mumu-store/app/models"
	"github.com/ipan97/mumu-store/config"
	"html/template"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Static("/public", "./public")
	r.HTMLRender = ginview.New(goview.Config{
		Root:         "./app/views",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     []string{},
		Funcs:        make(template.FuncMap),
		DisableCache: false,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})

	dbConfig := config.Config{
		Database:      config.NewPostgreSQLConfig(),
		IsDevelopment: true,
	}
	db, _ := dbConfig.GetDB()
	db.AutoMigrate(models.User{}, models.Category{}, models.Brand{}, models.Product{})

	// Web Route
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Mumu Store",
		})
	})

	return r
}
