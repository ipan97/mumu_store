package app

import (
	"github.com/ipan97/mumu-store/app/controllers/admin"
	"github.com/ipan97/mumu-store/app/controllers/api"
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

func (di *DependencyInjection) InjectCategoryController() *api.CategoryController {
	categoryRepository := repositories.NewCategoryRepository(di.db)
	categoryMapper := mapper.NewCategoryMapper()
	categoryService := services.NewCategoryService(categoryRepository, categoryMapper)
	return api.NewCategoryController(categoryService)
}

func (di *DependencyInjection) InjectBrandController() *api.BrandController {
	brandRepository := repositories.NewBrandRepository(di.db)
	brandMapper := mapper.NewBrandMapper()
	brandService := services.NewBrandService(brandRepository, brandMapper)
	return api.NewBrandController(brandService)
}

func (di *DependencyInjection) InjectDashboardController() *admin.DashboardController {
	return admin.NewDashboardController()
}
