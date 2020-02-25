package services

import (
	"github.com/ipan97/mumu-store/app/repositories"
	"github.com/ipan97/mumu-store/app/services/dto"
	"github.com/ipan97/mumu-store/app/services/mapper"
)

type BrandService struct {
	brandRepository repositories.BrandRepository
	brandMapper     *mapper.BrandMapper
}

func NewBrandService(brandRepository repositories.BrandRepository, brandMapper *mapper.BrandMapper) *BrandService {
	return &BrandService{brandRepository: brandRepository, brandMapper: brandMapper}
}

func (s *BrandService) FindAll() (*[]dto.BrandDto, error) {
	brands, err := s.brandRepository.FindAll()
	return s.brandMapper.ToDtos(&brands), err
}

func (s *BrandService) FindByID(id uint) (*dto.BrandDto, error) {
	brand, err := s.brandRepository.FindById(id)
	return s.brandMapper.ToDto(&brand), err
}
