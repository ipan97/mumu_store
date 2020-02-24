package repositories

import (
	"github.com/ipan97/mumu-store/app/models"
	"github.com/jinzhu/gorm"
)

type BrandRepository interface {
	FindAll() ([]models.Brand, error)
	FindById(uint) (models.Brand, error)
	Save(models.Brand) (models.Brand, error)
	Update(models.Brand) error
	Delete(models.Brand) error
	Count() (uint, error)
}

type brandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepository{db: db}
}

func (r brandRepository) FindAll() ([]models.Brand, error) {
	var brands []models.Brand
	err := r.db.Find(&brands).Error
	return brands, err
}

func (r brandRepository) FindById(id uint) (models.Brand, error) {
	var brand models.Brand
	err := r.db.First(&brand, id).Error
	return brand, err
}

func (r brandRepository) Save(brand models.Brand) (models.Brand, error) {
	err := r.db.Create(&brand).Error
	return brand, err
}

func (r brandRepository) Update(brand models.Brand) error {
	return r.db.UpdateColumns(&brand).Error
}

func (r brandRepository) Delete(brand models.Brand) error {
	return r.db.Delete(&brand).Error
}

func (r brandRepository) Count() (uint, error) {
	var count uint
	err := r.db.Model(&models.Brand{}).Count(&count).Error
	return count, err
}
