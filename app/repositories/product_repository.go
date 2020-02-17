package repositories

import "github.com/jinzhu/gorm"

type ProductRepository struct {
	db *gorm.DB
}
