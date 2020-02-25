package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ipan97/mumu-store/routers"
	"github.com/jinzhu/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	routers.Initialize(r, db)
	return r
}
