package services

import "github.com/ipan97/mumu-store/app/repositories"

type CategoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(categoryRepository repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository}
}

func (categoryService *CategoryService) FindAllCategories() {

}
