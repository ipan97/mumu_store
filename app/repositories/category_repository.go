package repositories

import (
	"github.com/ipan97/mumu-store/app/models"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func (r *CategoryRepository) FindAll() (*[]models.Category, error) {
	var categories []models.Category
	if err := r.db.Model(models.Category{}).Find(&categories).Error; err != nil {
		return nil, err
	} else {
		return &categories, nil
	}
}
