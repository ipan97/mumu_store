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

func (categoryRepository *CategoryRepository) FindAll() (*[]models.Category, error) {
	var categories []models.Category
	return &categories, categoryRepository.db.Model(models.Category{}).Find(&categories).Error
}

func (categoryRepository *CategoryRepository) FindById(id uint) (models.Category, error) {
	var category models.Category
	return category, categoryRepository.db.Model(models.Category{}).Take(&category, id).Error
}

func (categoryRepository *CategoryRepository) Save(category *models.Category) error {
	return categoryRepository.db.Model(models.Category{}).Save(&category).Error
}

func (categoryRepository *CategoryRepository) Update(category *models.Category) {

}

func (categoryRepository *CategoryRepository) Delete(id uint) {

}
