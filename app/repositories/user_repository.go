package repositories

import "github.com/jinzhu/gorm"

type UserRepository struct {
	db *gorm.DB
}
