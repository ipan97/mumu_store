package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ipan97/mumu-store/app/services"
)

type BrandController struct {
	brandService *services.BrandService
}

func NewBrandController(brandService *services.BrandService) *BrandController {
	return &BrandController{brandService: brandService}
}

func (bc *BrandController) FindAllBrands(ctx *gin.Context) {
	brands, err := bc.brandService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, brands)
	}
}

func (bc *BrandController) FindBrandByID(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.ParseInt(paramID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	brands, err := bc.brandService.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, brands)
		return
	}

}
