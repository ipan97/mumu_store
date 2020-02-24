package app

import (
	"github.com/ipan97/mumu-store/app/controllers"
	"github.com/ipan97/mumu-store/app/repositories"
	"github.com/ipan97/mumu-store/app/services"
	"github.com/ipan97/mumu-store/app/services/mapper"
	"github.com/jinzhu/gorm"
)

type DependencyInjection struct {
	db *gorm.DB
}

func NewDependencyInjection(db *gorm.DB) *DependencyInjection {
	return &DependencyInjection{db: db}
}

func (di *DependencyInjection) InjectCategoryController() *controllers.CategoryController {
	categoryRepository := repositories.NewCategoryRepository(di.db)
	categoryMapper := mapper.NewCategoryMapper()
	categoryService := services.NewCategoryService(categoryRepository, categoryMapper)
	categoryController := controllers.NewCategoryController(categoryService)
	return categoryController
}
