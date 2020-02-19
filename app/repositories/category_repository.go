package repositories

import (
	"github.com/ipan97/mumu-store/app/models"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) FindAll() (*[]models.Category, error) {
	var categories []models.Category
	if err := r.db.Model(models.Category{}).Find(&categories).Error; err != nil {
		return nil, err
	} else {
		return &categories, nil
	}
}
