package mapper

import (
	"github.com/ipan97/mumu-store/app/models"
	"github.com/ipan97/mumu-store/app/services/dto"
	"github.com/jinzhu/gorm"
)

type CategoryMapper struct {
}

func NewCategoryMapper() *CategoryMapper {
	return &CategoryMapper{}
}

func (categoryMapper *CategoryMapper) ToCategoryDto(category *models.Category) *dto.CategoryDto {
	if category == nil {
		return nil
	} else {
		return &dto.CategoryDto{
			ID:     category.ID,
			Name:   category.Name,
			Cover:  category.Cover,
			Slug:   category.Slug,
			Status: category.Status,
		}
	}
}

func (categoryMapper *CategoryMapper) ToCategory(categoryDto *dto.CategoryDto) *models.Category {
	if categoryDto == nil {
		return nil
	} else {
		return &models.Category{
			Model: gorm.Model{
				ID: categoryDto.ID,
			},
			Name:        categoryDto.Name,
			Cover:       categoryDto.Cover,
			Description: categoryDto.Description,
			Slug:        categoryDto.Slug,
			Status:      categoryDto.Status,
		}
	}
}

func (categoryMapper *CategoryMapper) ToCategoryDtos(categories *[]models.Category) *[]dto.CategoryDto {
	if categories == nil {
		return nil
	} else {
		categoryDtos := make([]dto.CategoryDto, len(*categories))
		for k, v := range *categories {
			categoryDtos[k] = *categoryMapper.ToCategoryDto(&v)
		}
		return &categoryDtos
	}
}

func (categoryMapper *CategoryMapper) ToCategories(categoryDtos *[]dto.CategoryDto) *[]models.Category {
	if categoryDtos == nil {
		return nil
	} else {
		categories := make([]models.Category, len(*categoryDtos))
		for k, v := range *categoryDtos {
			categories[k] = *categoryMapper.ToCategory(&v)
		}
		return &categories
	}
}
