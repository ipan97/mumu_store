package server

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/ipan97/mumu-store/routers"
	"github.com/jinzhu/gorm"
	"html/template"
)

func Setup(db *gorm.DB) *gin.Engine {
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
	routers.Initialize(r, db)
	r.Run()
	return r
}
