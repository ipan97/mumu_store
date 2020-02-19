package services

import (
	"github.com/ipan97/mumu-store/app/repositories"
	"github.com/ipan97/mumu-store/app/services/dto"
	"github.com/ipan97/mumu-store/app/services/mapper"
)

type CategoryService struct {
	categoryRepository *repositories.CategoryRepository
	categoryMapper     *mapper.CategoryMapper
}

func NewCategoryService(categoryRepository *repositories.CategoryRepository, categoryMapper *mapper.CategoryMapper) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository, categoryMapper: categoryMapper}
}

func (categoryService *CategoryService) FindAllCategories() (*[]dto.CategoryDto, error) {
	categories, err := categoryService.categoryRepository.FindAll()
	return categoryService.categoryMapper.ToCategoryDtos(categories), err
}
