package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ipan97/mumu-store/app/services"
	"net/http"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (categoryController *CategoryController) FindAllCategories(ctx *gin.Context) {
	categories, err := categoryController.categoryService.FindAllCategories()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, &categories)
	}
}
