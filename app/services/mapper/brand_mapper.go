package mapper

import (
	"github.com/ipan97/mumu-store/app/models"
	"github.com/ipan97/mumu-store/app/services/dto"
	"github.com/jinzhu/gorm"
)

type BrandMapper struct {
}

func NewBrandMapper() *BrandMapper {
	return &BrandMapper{}
}

func (m *BrandMapper) ToDto(brand *models.Brand) *dto.BrandDto {
	if brand == nil {
		return nil
	}

	return &dto.BrandDto{
		ID:        brand.ID,
		Name:      brand.Name,
		CreatedAt: brand.CreatedAt,
		UpdatedAt: brand.UpdatedAt,
		DeletedAt: brand.DeletedAt,
	}
}

func (m *BrandMapper) ToModel(brand *dto.BrandDto) *models.Brand {
	if brand == nil {
		return nil
	}

	return &models.Brand{
		Model: gorm.Model{
			ID: brand.ID,
		},
		Name: brand.Name,
	}
}

func (m *BrandMapper) ToDtos(brands *[]models.Brand) *[]dto.BrandDto {
	if brands == nil {
		return nil
	}
	brandDtos := make([]dto.BrandDto, len(*brands))
	for k, v := range *brands {
		brandDtos[k] = *m.ToDto(&v)
	}
	return &brandDtos
}

func (m *BrandMapper) ToModels(brandDtos *[]dto.BrandDto) *[]models.Brand {
	if brandDtos == nil {
		return nil
	}
	brands := make([]models.Brand, len(*brandDtos))
	for k, v := range *brandDtos {
		brands[k] = *m.ToModel(&v)
	}
	return &brands
}
